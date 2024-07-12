package templates

import (
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
)

func Run(app *fiber.App) {
	// Route to render the login template
	app.Get("/login", func(c fiber.Ctx) error {
		nonce := "" //c.Locals("nonce")
		return c.Render("login", fiber.Map{
			"Nonce": nonce,
		})
	})
	app.Get("/register", func(c fiber.Ctx) error {
		nonce := "" // c.Locals("nonce").(string)
		return c.Render("register", fiber.Map{
			"Nonce": nonce,
		})
	})

	app.Get("/profile", profile, auth.IsAuthenticated)

}

func profile(c fiber.Ctx) error {
	username, err := auth.GetUsername(c)
	if err != nil {
		log.Println("Profile page was visited without authentication.")
		return util.InternalServerError(c)
	}
	return c.Render("profile", fiber.Map{
		"Username":  username,
		"TotalBans": 0,
	})

}
