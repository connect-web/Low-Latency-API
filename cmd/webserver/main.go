package main

import (
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"log"
	"os"
	"time"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // or compress.LevelBestCompression
	}))

	app.Use(limiter.New(limiter.Config{
		Next: func(c fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c fiber.Ctx) error {
			c.Status(fiber.StatusTooManyRequests)
			return c.JSON(fiber.Map{"Error": "Rate limited try again later."})
		},
	}))

	// Define a route for the GET method on the root path '/'
	//app.Get("/api/ratio", api.GetPlayersByRatioHandler)
	//app.Get("/api/experience", api.GetPlayersByExperienceHandler)
	//app.Get("/api/levels", api.GetPlayersByLevelHandler)

	//app.Get("/api/users", api.GetSimplePlayerFromName)

	app.Get("/api/find-bots", api.GetPlayerFromSkills)

	app.Static("/", "../../site/", fiber.Static{
		CacheDuration: 30 * time.Minute,
		Compress:      true,
		Index:         "home.html",
	})

	envVar := os.Getenv("siteonline")
	certDirectory := os.Getenv("certDir")

	if envVar == "site" {
		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
		log.Fatal(app.Listen(":4050"))
	}

}
