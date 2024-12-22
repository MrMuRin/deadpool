package http

import (
	"deadpool/adapters/auth"
	"deadpool/core/services"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func (h *AuthHandler) GoogleLogin(c *fiber.Ctx) error {
	redirectURL := c.Query("redirectURL", "http://localhost:5174/menu") // ค่า Default Redirect
	state := base64.StdEncoding.EncodeToString([]byte(redirectURL))

	url := auth.GoogleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func (h *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No code in request"})
	}

	token, err := auth.GoogleOauthConfig.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to exchange token"})
	}

	userInfo, err := auth.GetUserInfo(token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}

	user, err := h.AuthService.LoginWithGoogle(userInfo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to login with Google"})
	}

	return c.JSON(user)
}
