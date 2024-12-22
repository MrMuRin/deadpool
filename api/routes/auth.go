package routes

import (
	"github.com/gofiber/fiber/v2"
)

func GoogleLogin(c *fiber.Ctx) error {
	// Logic สำหรับ Redirect ไปยัง Google OAuth
	return c.Redirect("https://accounts.google.com/o/oauth2/v2/auth?client_id=your_client_id")
}

func GoogleCallback(c *fiber.Ctx) error {
	// รับ Authorization Code จาก Google
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No code in the request",
		})
	}

	// Logic สำหรับแลกเปลี่ยน Code เป็น Token และดึงข้อมูลผู้ใช้
	// ส่ง Response กลับไป
	return c.JSON(fiber.Map{
		"message": "Google callback successful",
	})
}

func ProtectedEndpoint(c *fiber.Ctx) error {
	// ตัวอย่าง Protected Route
	return c.JSON(fiber.Map{
		"message": "This is a protected endpoint",
	})
}
