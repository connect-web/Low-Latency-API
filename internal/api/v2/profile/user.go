package profile

import (
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

func GetProfile(c fiber.Ctx) error {
	// method does not serve any purpose currently.

	username, err := auth.GetUsername(c)
	if err != nil || username == "" {
		return util.InternalServerError(c)
	}

	/*
		profile, profileFetchErr := user.FetchOrCreateProfile(username)
		if profileFetchErr != nil {
			return util.InternalServerError(c)
		}

		profile.Username = username

	*/

	return c.JSON(model.ProfileStats{Username: username})
}
