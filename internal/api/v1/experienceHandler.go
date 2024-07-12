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

func GetPlayersByExperienceHandler(c fiber.Ctx) error {
	experienceThreshold, err := extractExperienceParams(c)
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

	query, params := db.BuildPlayersByExperienceQuery(experienceThreshold)
	fmt.Println(query)
	fmt.Println(params)

	players, err := client.QueryDBSimplePlayers(query, params, db.HandleSimplePlayerRowExperience)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

func extractExperienceParams(c fiber.Ctx) (map[string]int64, error) {
	thresholds := map[string]int64{}

	queryParams := c.Context().QueryArgs()

	// Iterate over all query parameters
	queryParams.VisitAll(func(key, value []byte) {
		skillName := string(key)
		_, exists := util.SkillsMap[skillName]
		if exists {
			experienceThreshold, err := strconv.ParseInt(string(value), 10, 64)
			if err == nil {
				thresholds[util.Title.String(skillName)] = experienceThreshold
			}
		}
	})

	if len(thresholds) == 0 {
		return thresholds, errors.New("missing required parameters")
	}
	return thresholds, nil
}
