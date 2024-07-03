package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/utils/v2"
	"time"
)

var userSessionStore = session.New(session.Config{})

func Main(app *fiber.App) {

	csrfMiddleware := csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUIDv4,
	})

	app.Post("/api/register", Register, csrfMiddleware)
	app.Post("/api/login", Login, csrfMiddleware)
	app.Get("/api/csrf", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"Valid": true})
	}, csrfMiddleware)

}
