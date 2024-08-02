package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Run(app *fiber.App) {
	// This is for setting up the global middleware required on all routes.

	// Logging middleware
	app.Use(logger.New(logger.Config{}))

	// Compression
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // or compress.LevelBestCompression
	}))

	// Use Helmet middleware with custom CSP
	/*
		app.Use(helmet.New(helmet.Config{
				ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'nonce-%s' https://hcaptcha.com https://*.hcaptcha.com; frame-src 'self' https://hcaptcha.com https://*.hcaptcha.com; style-src 'self' 'nonce-%s' https://hcaptcha.com https://*.hcaptcha.com; img-src 'self' https://*.hcaptcha.com; connect-src 'self' https://hcaptcha.com https://*.hcaptcha.com",
			}))
	*/
}
