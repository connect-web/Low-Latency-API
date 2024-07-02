package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Run(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // or compress.LevelBestCompression
	}))

	// Logging middleware
	app.Use(logger.New(logger.Config{}))

	RateLimit(app)

	RewriteEngine(app)

}
