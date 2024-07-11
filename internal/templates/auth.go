package templates

import (
	"github.com/gofiber/fiber/v3"
)

func Run(app *fiber.App) {
	// Route to render the login template
	app.Get("/login", func(c fiber.Ctx) error {
		nonce := c.Locals("nonce").(string)
		return c.Render("login", fiber.Map{
			"Nonce": nonce,
		})
	})

	// Route to render the register template
	app.Get("/register", func(c fiber.Ctx) error {
		nonce := c.Locals("nonce").(string)
		return c.Render("register", fiber.Map{
			"Nonce": nonce,
		})
	})
}
