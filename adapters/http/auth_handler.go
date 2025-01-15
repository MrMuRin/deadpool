package http

import (
	"deadpool/adapters/auth"
	"deadpool/core/services"
	"deadpool/utils"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	AuthService *services.AuthService
	GoogleAuth  *auth.GoogleAuth
}

func NewAuthHandler(authService *services.AuthService, googleAuth *auth.GoogleAuth) *AuthHandler {
    return &AuthHandler{
        AuthService: authService,
        GoogleAuth:  googleAuth,
    }
}

func (h *AuthHandler) GoogleLogin(c *fiber.Ctx) error {
	redirectURL := c.Query("redirectURL", "http://localhost:5173/menu") // ค่า Default Redirect
	state := base64.StdEncoding.EncodeToString([]byte(redirectURL))

	authCodeURL := h.GoogleAuth.Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(authCodeURL)
}

func (h *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
    code := c.Query("code")
    state := c.Query("state")

    if code == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No code in request"})
    }

    // ถอดรหัส state เพื่อดึง redirectURL
    decodedState, err := base64.StdEncoding.DecodeString(state)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid state"})
    }
    redirectURL := string(decodedState)

    token, err := h.GoogleAuth.Config.Exchange(c.Context(), code)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
    }

    userInfo, err := h.GoogleAuth.GetUserInfo(token)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
    }

    user, err := h.AuthService.LoginWithGoogle(userInfo)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login with Google"})
    }

    jwtToken, err := utils.GenerateJWT(user.ID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
    }

    c.Cookie(&fiber.Cookie{
        Name:     "authToken",
        Value:    jwtToken,
        Expires:  time.Now().Add(24 * time.Hour),
        HTTPOnly: true,
        Secure:   false,
        SameSite: "Strict",
    })

    // Redirect ผู้ใช้ไปยัง redirectURL พร้อมส่ง Token
    return c.Redirect(fmt.Sprintf("%s?token=%s", redirectURL, jwtToken))
}

