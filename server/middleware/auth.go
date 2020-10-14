package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// Protected protect routes
func Protected() fiber.Handler {

	return func(c *fiber.Ctx) error {
		auth := c.Cookies("user")
		if auth != "" {
			return c.Next()
		}
		return c.Redirect("/login")

	}
}
