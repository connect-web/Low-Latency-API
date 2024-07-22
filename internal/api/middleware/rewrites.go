package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"strings"
)

func RewriteEngine(app *fiber.App) {
	app.Use(Rewrites)
}

func Rewrites(c fiber.Ctx) error {
	path := c.Path()
	avoidContainingPaths := []string{
		"api",
		"resources",
		"favicon",
	}

	fmt.Println(path)
	for _, avoidStr := range avoidContainingPaths {
		if strings.Contains(strings.ToLower(path), avoidStr) {
			return c.Next()
		}
	}

	// Applying the rewrite rules directly
	rules := map[string]string{
		"/":  "/home.html",
		"/*": "/$1.html",
	}

	for pattern, replacement := range rules {
		if pattern == "/*" {
			// Handle dynamic paths
			if path != "/" && !strings.HasSuffix(path, ".html") {
				c.Path(strings.TrimSuffix(path, "/") + ".html")
			}
		} else if path == pattern {
			// Handle static paths
			c.Path(replacement)
		}
	}

	return c.Next()
}
