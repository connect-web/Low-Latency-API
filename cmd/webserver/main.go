package main

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/connect-web/Low-Latency-API/internal/auth"
	"github.com/connect-web/Low-Latency-API/internal/middleware"
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
	// app.Use(helmet.New(), )

	middleware.Run(app)

	API(app)
	StaticFiles(app)

	auth.Main(app)
	//auth.UnauthenticatedSessionSetup(app)

	run(app)

}

func run(app *fiber.App) {
	if front_end {
		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
		log.Fatal(app.Listen(":4050"))
	}
}

func API(app *fiber.App) {
	app.Get("/api/find-skill-bots", api.GetPlayerFromSkills)
	app.Get("/api/find-minigame-bots", api.GetPlayerFromMinigames)
	app.Get("/api/minigame-bots-hiscore", api.PlayerMinigameHiscores)
	app.Get("/api/minigame-bots-listing", api.PlayerMinigameListing)

	// Define a route for the GET method on the root path '/'
	//app.Get("/api/ratio", api.GetPlayersByRatioHandler)
	//app.Get("/api/experience", api.GetPlayersByExperienceHandler)
	//app.Get("/api/levels", api.GetPlayersByLevelHandler)
	//app.Get("/api/users", api.GetSimplePlayerFromName)
}

func StaticFiles(app *fiber.App) {
	staticType := fiber.Static{
		Index: "home",
	}
	if front_end {
		staticType.CacheDuration = 30 * time.Minute
		staticType.Compress = true
	}
	app.Static("/", "../../site/")
}
