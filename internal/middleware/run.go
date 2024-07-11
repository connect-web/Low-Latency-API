package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// GenerateNonce generates a random nonce for each request
func GenerateNonce() string {
	nonce := make([]byte, 16)
	_, err := rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(nonce)
}

func setCSP() fiber.Handler {
	return func(c fiber.Ctx) error {
		//nonce := GenerateNonce()
		//c.Locals("nonce", nonce)
		c.Set("Content-Security-Policy", fmt.Sprintf(
			"default-src 'self'; "+
				"script-src 'self' https://hcaptcha.com https://*.hcaptcha.com; "+
				"frame-src 'self' https://hcaptcha.com https://*.hcaptcha.com; "+
				"style-src 'self' https://hcaptcha.com https://*.hcaptcha.com; "+
				"img-src 'self' https://*.hcaptcha.com; "+
				"connect-src 'self' https://hcaptcha.com https://*.hcaptcha.com",
		))
		return c.Next()
	}
}

func Run(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // or compress.LevelBestCompression
	}))

	// Use Helmet middleware with custom CSP
	/*
		app.Use(helmet.New(helmet.Config{
				ContentSecurityPolicy: "default-src 'self'; script-src 'self' 'nonce-%s' https://hcaptcha.com https://*.hcaptcha.com; frame-src 'self' https://hcaptcha.com https://*.hcaptcha.com; style-src 'self' 'nonce-%s' https://hcaptcha.com https://*.hcaptcha.com; img-src 'self' https://*.hcaptcha.com; connect-src 'self' https://hcaptcha.com https://*.hcaptcha.com",
			}))
	*/

	app.Use(setCSP())

	// Logging middleware
	app.Use(logger.New(logger.Config{}))

	RateLimit(app)

	RewriteEngine(app)

}
