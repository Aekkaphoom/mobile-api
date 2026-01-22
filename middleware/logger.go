package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		log.Printf("[%d] %s %s %s",
			c.Response().StatusCode(),
			c.Method(),
			c.Path(),
			time.Since(start),
		)
		return err
	}
}
