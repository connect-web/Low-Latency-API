package main

import (
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"log"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // or compress.LevelBestCompression
	}))

	// Define a route for the GET method on the root path '/'
	app.Get("/ratio", api.GetPlayersByRatioHandler)

	// Start the server on port 3000
	log.Fatal(app.Listen(":4050"))
}
