package main

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/auth"
	"github.com/connect-web/Low-Latency-API/internal/middleware"
	"github.com/connect-web/Low-Latency-API/internal/protectedApis"
	"github.com/gofiber/fiber/v3"
	"log"
	"os"
	"time"
)

var (
	envVar        = os.Getenv("siteonline")
	certDirectory = os.Getenv("certDir")
	front_end     = envVar == "site" // True if front_end , False if local development
)

func main() {
	fmt.Printf("Front end mode = %v\n", front_end)
	// Initialize a new Fiber app
	app := fiber.New()
	// Use middlewares for each route

	middleware.Run(app) // Setup middleware

	// Setup Static files routes
	staticType := fiber.Static{
		Index: "home",
	}
	if front_end {
		staticType.CacheDuration = 30 * time.Minute
		staticType.Compress = true
	}
	app.Static("/", "../../site/")

	auth.Setup(app) // Setup Register, Login Routes

	protectedApis.Setup(app)

	if front_end {
		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
		log.Fatal(app.Listen(":4050"))
	}

}
