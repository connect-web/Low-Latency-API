package middleware

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/util"
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

func NonceAndCSPMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Generate the nonce
		fmt.Println("Creating nonce.")
		nonce, err := util.GenerateNonce()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
		fmt.Printf("Nonce set: '%s'\n", nonce)
		c.Locals("nonce", nonce)

		// Set the CSP headers
		fmt.Println("Fetching nonce.")
		c.Set("Content-Security-Policy", fmt.Sprintf(
			"default-src 'self'; "+
				"script-src 'self' 'nonce-%[1]s' https://*.hcaptcha.com https://cdnjs.cloudflare.com https://cdn.jsdelivr.net; "+
				"style-src 'self' 'nonce-%[1]s' https://*.hcaptcha.com https://fonts.googleapis.com https://cdn.jsdelivr.net; "+
				"img-src 'self' https://*.hcaptcha.com https://cdnjs.cloudflare.com; "+
				"font-src 'self' https://fonts.googleapis.com https://fonts.gstatic.com; "+
				"frame-src 'self' https://hcaptcha.com https://*.hcaptcha.com; "+
				"connect-src 'self' https://hcaptcha.com https://*.hcaptcha.com",
			nonce,
		))
		return c.Next()
	}
}
