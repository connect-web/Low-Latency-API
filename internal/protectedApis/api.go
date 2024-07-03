package protectedApis

import (
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/connect-web/Low-Latency-API/internal/auth"
	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	app.Get("/api/get-name", GetUsername, Protected)

	app.Get("/api/find-skill-bots", api.GetPlayerFromSkills, Protected)
	app.Get("/api/find-minigame-bots", api.GetPlayerFromMinigames, Protected)
	app.Get("/api/minigame-bots-hiscore", api.PlayerMinigameHiscores, Protected)
	app.Get("/api/minigame-bots-listing", api.PlayerMinigameListing, Protected)

}

func GetUsername(c fiber.Ctx) error {
	sess, err := auth.UserSessionStore.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session retrieval error"})
	}

	username := sess.Get("username")
	if username == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Username not found in session"})
	}

	return c.JSON(fiber.Map{"message": username})
}
