package auth

import (
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

func Protected(c fiber.Ctx) error {
	username, err := GetUsername(c)
	if err != nil {
		return util.Unauthorized(c)
	}
	if username == "" {
		return util.Unauthorized(c)
	}
	return c.Next()
}

func IsAuthenticated(c fiber.Ctx) error {
	// Get the session from the context
	sess, err := UserSessionStore.Get(c)
	if err != nil {
		// Handle session retrieval error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session retrieval error"})
	}

	// Check if the user is authenticated
	if sess.Get("username") == nil {
		return c.Redirect().To("/login")
	}

	// Proceed if authenticated
	return c.Next()
}

func GetUsername(c fiber.Ctx) (string, error) {
	sess, err := UserSessionStore.Get(c)
	if err != nil {
		// Handle session retrieval error
		return "", err
	}

	// Check if the user is authenticated
	if sess.Get("username") == nil {
		return "", errors.New("Unauthenticated")
	}
	return sess.Get("username").(string), nil
}
