package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/rewrite"
	"strings"
)

func RewriteEngine(app *fiber.App) {
	app.Use(func(c fiber.Ctx) error {
		path := c.Path()
		avoid_paths := []string{
			"api",
			"resources",
		}
		for _, avoid_str := range avoid_paths {
			if strings.Contains(strings.ToLower(path), avoid_str) {
				return c.Next()
			}
		}

		rewriteMiddleware := rewrite.New(rewrite.Config{
			Rules: map[string]string{
				"/":  "/home.html",
				"/*": "/$1.html",
			},
		})
		return rewriteMiddleware(c)

	})
}
