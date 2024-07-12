package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/rewrite"
	"strings"
)

func RewriteEngine(app *fiber.App) {
	app.Use(Rewrites)
}

func Rewrites(c fiber.Ctx) error {
	path := c.Path()
	avoid_containing_paths := []string{
		"api",
		"resources",
	}

	avoid_paths := []string{
		"/login",
		"/register",
	}

	fmt.Println(path)
	for _, avoid_str := range avoid_containing_paths {
		if strings.Contains(strings.ToLower(path), avoid_str) {
			return c.Next()
		}
	}

	for _, avoidable := range avoid_paths {
		if strings.ToLower(path) == avoidable {
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

}
