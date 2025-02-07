package middlewares

import (
	"deadpool/core/services"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	AuthService *services.AuthService
}

func NewAuthMiddleware(authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		AuthService: authService,
	}
}

func (m *AuthMiddleware) Handle(c *fiber.Ctx) error {
	token := c.Cookies("authToken")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID, err := m.AuthService.ValidateToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}

	c.Locals("userID", userID)
	return c.Next()
}
