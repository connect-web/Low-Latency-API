package v1

import (
	"errors"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"log"
	"strings"
)

func GetSimplePlayerFromName(c fiber.Ctx) error {
	usernameMap, err := extractNameParams(c)
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

	query, params := db.BuildUsersQuery(usernameMap)
	fmt.Println(query)
	fmt.Println(params)

	players, err := client.QueryDBPlayers(query, util.ConvertStringArrayToInterfaceArray(params), db.HandlePlayerSkillsRow)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

func extractNameParams(c fiber.Ctx) (map[string]struct{}, error) {
	names := c.Query("names") // use Query instead of Params if it's a query parameter
	extractedNames := map[string]struct{}{}

	if names == "" {
		return nil, errors.New("missing required parameter: names")
	}
	for _, username := range strings.Split(names, ",") {
		if 12 < len(username) {
			continue // no name is above 12 chars...
		}
		extractedNames[username] = struct{}{}
	}

	if len(extractedNames) == 0 {
		return nil, errors.New("No valid usernames.")
	}

	if 120 < len(extractedNames) {
		return nil, errors.New("Exceeded limit 120 names per request.")
	}

	return extractedNames, nil
}
