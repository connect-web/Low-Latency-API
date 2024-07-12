package main

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/connect-web/Low-Latency-API/internal/api/middleware"
	"github.com/connect-web/Low-Latency-API/internal/templates"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
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

	engine := html.New("../../templates", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	// Setup middleware for all Routers
	middleware.Run(app)

	// Setup API's
	api_routes := app.Group("/api")
	api.CreateRouter(api_routes)

	// Setup Static files
	RegisterStatic(app)

	templates.Run(app)

	// Display environment, Dev / Server
	fmt.Printf("Front end mode = %v\n", front_end)

	if front_end {
		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
		log.Fatal(app.Listen(":4050"))
	}

}

func RegisterStatic(app *fiber.App) {
	staticType := fiber.Static{Index: "home"}
	if front_end {
		// Only cache and Compress outside of development.
		staticType.CacheDuration = 30 * time.Minute
		staticType.Compress = true
	}

	middleware.RewriteEngine(app) // apply Rewrite engine to
	app.Static("/", "../../site/", staticType)
}
