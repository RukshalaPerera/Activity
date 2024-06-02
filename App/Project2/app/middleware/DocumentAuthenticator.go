package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

func DocumentAuthenticate(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "unauthorized",
		})
	}
	role := strings.TrimPrefix(token, "Bearer ")
	c.Locals("userRole", role)

	return c.Next()
}
