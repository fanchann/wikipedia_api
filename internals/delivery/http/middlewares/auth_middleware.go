package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func NewAuthMiddlewares(log *logrus.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("api-key", "NOT_FOUND")
		if apiKey == "sup3rs3cret" {
			log.Infof("Log from %s", c.IP())
			return c.Next()
		}
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.ErrUnauthorized)
	}
}
