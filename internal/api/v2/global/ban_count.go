package global

import (
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/model"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

func GetBanCount(c fiber.Ctx) error {
	username, err := auth.GetUsername(c)
	if err != nil || username == "" {
		return util.InternalServerError(c)
	}
	return c.JSON(fiber.Map{
		"TotalBans": model.TotalBans,
	})
}
