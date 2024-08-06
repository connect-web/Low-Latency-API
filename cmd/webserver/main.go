package main

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/api/middleware"
	"github.com/connect-web/Low-Latency-API/internal/api/templates"
	"github.com/connect-web/Low-Latency-API/internal/db/globalStats"
	cache "github.com/connect-web/low-latency-cache-controller/wrapper"
	json "github.com/goccy/go-json"
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

func listRoutes(app *fiber.App) {
	fmt.Println("Defined routes:")
	for _, routes := range app.Stack() {
		for _, route := range routes {
			fmt.Printf("%s %s\n", route.Method, route.Path)
		}
	}
}

func main() {
	go globalstats.GlobalStatsWorker() // starts a worker that periodically updates the global ban count statistics

	engine := html.New("../../templates", ".html")

	// Create a new Fiber app with the HTML engine
	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	auth.CreateSessionStore()
	// Setup middleware for all Routers
	middleware.Run(app)
	// templates.Run(app) // templates not used yet

	// Setup API's
	api_routes := app.Group("/api")

	api.CreateRouter(api_routes)

	// Setup Static files
	templates.CreateTemplates(app)
	RegisterStatic(app)

	// Display environment, Dev / Server
	fmt.Printf("Front end mode = %v\n", front_end)
	listRoutes(app)
	if front_end {
		go cache.StartUp("https://low-latency.co.uk")
		go cache.RefreshCacheHourly("https://low-latency.co.uk")

		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
		go cache.StartUp("http://127.0.0.1:4050")
		go cache.RefreshCacheHourly("http://127.0.0.1:4050")

		log.Fatal(app.Listen(":4050"))
	}

}

func RegisterStatic(app *fiber.App) {
	staticType := fiber.Static{Index: "home"}
	if front_end {
		// Only cache and Compress outside of development.
		staticType.CacheDuration = 24 * 7 * time.Hour
		staticType.Compress = true
	}

	middleware.RewriteEngine(app) // apply Rewrite engine to
	app.Static("/", "../../site/", staticType)
}
