package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
    // Example middleware logic
    return c.Next()
}
