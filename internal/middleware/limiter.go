package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"strings"
	"time"
)

func RateLimit(app *fiber.App) {
	app.Use(limiter.New(limiter.Config{
		Next: func(c fiber.Ctx) bool {
			return c.IP() == "127.0.0.1" || !strings.Contains(c.Path(), "api")
		},
		Max:        5,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c fiber.Ctx) error {
			c.Status(fiber.StatusTooManyRequests)
			return c.JSON(fiber.Map{"Error": "Rate limited try again later."})
		},
	}))
}
