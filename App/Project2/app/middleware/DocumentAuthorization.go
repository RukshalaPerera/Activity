package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func DocumentAuthorization(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole := c.Locals("userRole").(string)

		for _, role := range allowedRoles {
			if userRole == role {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"success": false,
			"message": "Forbidden",
		})
	}
}
