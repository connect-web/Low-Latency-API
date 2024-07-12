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
	// Custom error handler
	ErrorHandler: func(c fiber.Ctx, err error) error {
		if err.Error() == "forbidden" {
			// Remove the CSRF cookie
			c.Cookie(&fiber.Cookie{
				Name:     "csrf_",
				Value:    "",
				Expires:  time.Now().Add(-1 * time.Hour),
				SameSite: "Lax",
			})
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "CSRF token is invalid or missing",
		})
	},
})

func CreateRouter(app fiber.Router) {
	app.Post("/register", Register)
	app.Post("/login", Login)
	app.Post("/logout", Logout)
	app.Get("/csrf", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"Valid": true})
	})

}

func createSession(username string, c fiber.Ctx) error {
	sess, err := UserSessionStore.Get(c)
	if err != nil {
		return err
	}
	sess.Set("username", username)
	return sess.Save()
}
