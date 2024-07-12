package v1

import (
	"errors"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
)

func GetPlayersByLevelHandler(c fiber.Ctx) error {
	LevelThreshold, err := extractLevelParams(c)
	if err != nil {
		return util.ErrorResponse(c, err.Error())
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

	query, params := db.BuildPlayersByLevelQuery(LevelThreshold)
	fmt.Println(query)
	fmt.Println(params)

	players, err := client.QueryDBSimplePlayers(query, params, db.HandleSimplePlayerRowLevel)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

func extractLevelParams(c fiber.Ctx) (map[string]int, error) {
	thresholds := map[string]int{}

	queryParams := c.Context().QueryArgs()

	// Iterate over all query parameters
	queryParams.VisitAll(func(key, value []byte) {
		skillName := string(key)
		_, exists := util.SkillsMap[skillName]
		if exists {
			LevelThreshold, err := strconv.ParseInt(string(value), 10, 32)
			if err == nil {
				thresholds[util.Title.String(skillName)] = int(LevelThreshold)
			}
		}
	})

	if len(thresholds) == 0 {
		return thresholds, errors.New("missing required parameters")
	}
	return thresholds, nil
}
