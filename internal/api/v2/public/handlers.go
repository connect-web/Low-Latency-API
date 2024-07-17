package public

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

/*
*** SKILLER TOPLISTS ***
 */
func GetSkillToplist(c fiber.Ctx) error {
	skillers, err := QuerySkillToplist()
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(skillers) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(skillers)
}

func GetSkillToplistUsers(c fiber.Ctx) error {
	string_skill := c.Query("skill-id")
	skillId, err := strconv.Atoi(string_skill)
	if err != nil {
		return util.InternalServerError(c)
	}
	fmt.Println(skillId)
	players, err := QuerySkillToplistUsers(skillId)
	if err != nil {
		fmt.Println(err.Error())
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}

/*
*** Boss & Minigame Toplists ***
 */

func GetMinigameToplist(c fiber.Ctx) error {
	minigamers, err := QueryMinigameToplist()
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(minigamers) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(minigamers)
}

func GetMinigameToplistUsers(c fiber.Ctx) error {
	string_skill := c.Query("minigame") // this should be converted to an ID

	valid := util.ValidMinigame(string_skill)
	if !valid {
		return util.InternalServerError(c)
	}

	players, err := QueryMinigameToplistUsers(string_skill)
	if err != nil {
		return util.InternalServerError(c)
	}

	if len(players) == 0 {
		return util.NoPlayersFound(c)
	}

	return c.JSON(players)
}
