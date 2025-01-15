package http

import (
	"deadpool/core/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
    return &UserHandler{
        UserService: userService,
    }
}

func (h *UserHandler) GetMe(c *fiber.Ctx) error {
    userID := c.Locals("userID").(uint)

    user, err := h.UserService.GetUserByID(userID)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user"})
    }

    return c.JSON(fiber.Map{
        "id":     user.ID,
        "name":   user.Name,
        "email":  user.Email,
        "avatar": user.Avatar,
    })
}