package api

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
	"net/url"
)

func PlayerMinigameHiscores(c fiber.Ctx) error {
	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return util.InternalServerError(c)
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	players, err := client.QueryMinigameHiscores()
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

func PlayerMinigameListing(c fiber.Ctx) error {
	minigame := c.Query("minigame")
	var err error
	minigame, err = url.QueryUnescape(minigame)

	if err != nil {
		// failed to decode query params.
		return util.InternalServerError(c)
	}

	if minigame == "" || !util.ValidMinigame(minigame) {
		fmt.Println(minigame)
		return util.ErrorResponse(c, "Missing minigame / boss / activity.")
	}

	client := db.NewDBClient()
	if connectErr := client.Connect(); connectErr != nil {
		log.Println(connectErr.Error())
		return util.InternalServerError(c)
	}
	defer func() {
		if closeConnErr := client.Close(); closeConnErr != nil {
			log.Printf("Failed to close database connection: %v", closeConnErr)
		}
	}()

	players, err := client.QueryMinigameListing(minigame, db.HandlePlayerRowMinigames)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}
