package protectedApis

import (
	"github.com/connect-web/Low-Latency-API/internal/auth"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

func Protected(c fiber.Ctx) error {
	sess, err := auth.UserSessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session retrieval error"})
	}
	if sess.Get("username") == nil {
		return util.Unauthorized(c)
	}
	return c.Next()
}
