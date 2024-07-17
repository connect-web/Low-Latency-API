package templates

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/db/user"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
)

func Run(app *fiber.App) {
	//AuthTemplates(app)
	UserTemplates(app)

}

func AuthTemplates(app *fiber.App) {
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
}

func UserTemplates(app *fiber.App) {
	// User api's
	app.Get("/profile", profile, auth.IsAuthenticated)
	app.Get("/search", search, auth.IsAuthenticated)

}

func profile(c fiber.Ctx) error {
	username, err := auth.GetUsername(c)

	if err != nil {
		log.Println("Profile page was visited without authentication.")
		return c.Redirect().To("/login")
	}

	userProfile, profileFetchErr := user.FetchOrCreateProfile(username)
	if profileFetchErr != nil {
		return util.InternalServerError(c)
	}
	fmt.Println(userProfile)
	return c.Render("profile", fiber.Map{
		"Username":         username,
		"BotsTracked":      userProfile.BotsTracked,
		"BotsBanned":       userProfile.BotsBanned,
		"BannedExperience": userProfile.BannedExperience,
		"PlayersAdded":     userProfile.PlayersAdded,

		"TotalBans": 0,
	})
}

func search(c fiber.Ctx) error {
	username, err := auth.GetUsername(c)
	if err != nil {
		log.Println("Search page was visited without authentication.")
		return c.Redirect().To("/login")
	}
	return c.Render("search", fiber.Map{
		"Username":  username,
		"TotalBans": 0,
	})
}
