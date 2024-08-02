package profile

import (
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/db/user"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

func GetProfile(c fiber.Ctx) error {
	username, err := auth.GetUsername(c)
	if err != nil || username == "" {
		return util.InternalServerError(c)
	}

	profile, profileFetchErr := user.FetchOrCreateProfile(username)
	if profileFetchErr != nil {
		return util.InternalServerError(c)
	}

	profile.Username = username

	return c.JSON(profile)
}
