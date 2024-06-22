package util

import "github.com/gofiber/fiber/v3"

func ErrorResponse(c fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Error": message,
	})
}

func NoPlayersFound(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"Error": "No players found",
	})
}

func InternalServerError(c fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Error": "Internal server error",
	})
}
