package api

import (
	"errors"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
	"strings"
)

func GetPlayerFromMinigames(c fiber.Ctx) error {
	selectedSkills, minigameThresholds, dailyXpThresholds, minLevels, maxLevels, err := extractMinigameBotFilterParams(c)
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

	query, params := db.BuildMinigameBotFinderQuery(selectedSkills, minigameThresholds, dailyXpThresholds, minLevels, maxLevels)

	players, err := client.QueryDBPlayers(query, params, db.HandlePlayerRowMinigames)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

func extractMinigameBotFilterParams(c fiber.Ctx) (map[string]struct{}, map[string]int, map[string]int, map[string]int, map[string]int, error) {
	queryParams := c.Context().QueryArgs()

	selectedSkills := map[string]struct{}{}

	minigameThreshold := map[string]int{}

	dailyXpThresholds := map[string]int{}
	minLevels := map[string]int{}
	maxLevels := map[string]int{}

	// Iterate over all query parameters
	queryParams.VisitAll(func(key, value []byte) {
		keyText := string(key)
		if !strings.Contains(keyText, "_") {
			_, exists := util.SkillsMap[keyText]
			if exists {
				selectedSkills[util.Title.String(keyText)] = struct{}{}
			}

			_, minigameExists := util.MinigamesMap[keyText]
			if minigameExists {
				val, intParseErr := util.StringToIntText(string(value))
				if intParseErr == nil {
					minigameThreshold[util.Title.String(keyText)] = val
				}

			}

			return
		}

		skillName := strings.Split(keyText, "_")[0]

		if strings.Contains(keyText, "_daily") {
			_, exists := util.SkillsMap[skillName]
			if exists {
				val, intParseErr := util.StringToIntText(string(value))
				if intParseErr == nil {
					dailyXpThresholds[util.Title.String(skillName)] = val
				}
			}
		}

		if strings.Contains(keyText, "_min_lvl") {
			_, exists := util.SkillsMap[skillName]
			if exists {
				number, err := strconv.Atoi(string(value))
				if err == nil {
					if validLevel(number) {
						minLevels[util.Title.String(skillName)] = number
					}

				}
			}
		}

		if strings.Contains(keyText, "_max_lvl") {
			_, exists := util.SkillsMap[skillName]
			if exists {
				number, err := strconv.Atoi(string(value))
				if err == nil {
					if validLevel(number) {
						maxLevels[util.Title.String(skillName)] = number
					}

				}
			}
		}
	})

	if len(selectedSkills) == 0 {
		return selectedSkills, minigameThreshold, dailyXpThresholds, minLevels, maxLevels, errors.New("No skills selected.")
	}

	return selectedSkills, minigameThreshold, dailyXpThresholds, minLevels, maxLevels, nil
}
