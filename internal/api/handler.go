package api

import (
	"errors"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
	"strconv"
	"strings"
)

func GetPlayersByRatioHandler(c fiber.Ctx) error {
	skillArray, ratioThreshold, err := extractRatioParams(c)
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

	query, params := db.BuildPlayersByRatioQuery(skillArray, ratioThreshold)
	fmt.Println(query)
	fmt.Println(params)
	players, err := client.QueryDBSimplePlayers(query, params, db.HandleSimplePlayerRow)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

func extractRatioParams(c fiber.Ctx) ([]string, float64, error) {
	skills := c.Query("skills") // use Query instead of Params if it's a query parameter
	ratio := c.Query("ratio")

	fmt.Println(skills)
	fmt.Println(ratio)

	if skills == "" || ratio == "" {
		return nil, 0, errors.New("missing required parameters: skills and/or ratio")
	}

	skillArray := util.ValidateSkills(strings.Split(skills, ","))
	if len(skillArray) == 0 {
		return nil, 0, errors.New("you entered no skills")
	}
	if len(skillArray) > 5 {
		return nil, 0, errors.New("you entered too many skills")
	}

	ratioThreshold, err := strconv.ParseFloat(ratio, 64)
	if err != nil {
		return nil, 0, errors.New("invalid ratio format")
	}

	return skillArray, ratioThreshold, nil
}
