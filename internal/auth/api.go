package auth

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/utils/v2"
	"time"
)

var UserSessionStore = session.New(session.Config{})

var CsrfMiddleware = csrf.New(csrf.Config{
	KeyLookup:      "header:X-Csrf-Token",
	CookieName:     "csrf_",
	CookieSameSite: "Lax",
	Expiration:     1 * time.Hour,
	KeyGenerator:   utils.UUIDv4,
})

func Setup(app *fiber.App) {

	app.Post("/api/register", Register, CsrfMiddleware)
	app.Post("/api/login", Login, CsrfMiddleware)
	app.Post("/api/logout", Logout, CsrfMiddleware)

	app.Get("/api/csrf", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"Valid": true})
	}, CsrfMiddleware)

}
