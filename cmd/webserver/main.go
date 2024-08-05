package main

import (
	"context"
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api"
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/api/middleware"
	"github.com/connect-web/Low-Latency-API/internal/api/templates"
	"github.com/connect-web/Low-Latency-API/internal/db"
	"github.com/connect-web/Low-Latency-API/internal/db/globalStats"
	"github.com/connect-web/storageself/postgres"
	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/template/html/v2"
	"github.com/jackc/pgx/v5/pgxpool"
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

	config, err := pgxpool.ParseConfig(db.GetUrl())

	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL %v\n", err)
	}

	config.MaxConns = 2

	pool, connectConfigErr := pgxpool.NewWithConfig(context.Background(), config)
	if connectConfigErr != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	defer pool.Close()

	store := postgres.New(
		postgres.Config{
			ConnectionURI: db.GetUrl(),
			Table:         "users.fiber_storage",
			Reset:         false,
			GCInterval:    1 * time.Minute,
		})
	auth.UserSessionStore = session.New(session.Config{Storage: store})

	engine := html.New("../../templates", ".html")

	// Create a new Fiber app with the HTML engine
	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

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
		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
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
