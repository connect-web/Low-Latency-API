package auth

import (
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/storageself/postgres"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/utils/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
	"log"
	"time"
)

var UserSessionStore *session.Store

func CreateSessionStore() {
	config, err := pgxpool.ParseConfig(db.GetUrl())

	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL %v\n", err)
	}

	config.MaxConns = 2

	pool, connectConfigErr := pgxpool.NewWithConfig(context.Background(), config)
	if connectConfigErr != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	store := postgres.New(
		postgres.Config{
			ConnectionURI: db.GetUrl(),
			Table:         "users.fiber_storage",
			Reset:         false,
			GCInterval:    1 * time.Minute,
		})
	UserSessionStore = session.New(session.Config{Storage: store})
}

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
