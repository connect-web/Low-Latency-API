package templates

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
)

var (
	TemplatePaths = map[string]string{
		// Key=Route path
		// Value=File path
		"/home":     "home",
		"/":         "home",
		"/login":    "login",
		"/register": "register",
		"/profile":  "profile",

		"/search/skills":    "search/skills",
		"/search/minigames": "search/minigames",
		"/ml/skills":        "ml/skills",
		"/ml/minigames":     "ml/minigames",
	}
	ALLOWED_DOMAINS = map[string]string{
		"script":  "https://*.hcaptcha.com https://cdnjs.cloudflare.com https://cdn.jsdelivr.net",
		"style":   "https://*.hcaptcha.com https://fonts.googleapis.com https://cdn.jsdelivr.net",
		"img":     "https://*.hcaptcha.com https://cdnjs.cloudflare.com",
		"font":    "https://fonts.googleapis.com https://fonts.gstatic.com  https://cdnjs.cloudflare.com",
		"frame":   "https://hcaptcha.com https://*.hcaptcha.com",
		"connect": "https://hcaptcha.com https://*.hcaptcha.com",
	}
)

func CreateTemplates(app *fiber.App) {
	for api_path, local_path := range TemplatePaths {
		//api_path = fmt.Sprintf("%s-test", api_path)
		app.Get(api_path, renderWithNonce(local_path), NonceAndCSPMiddleware())
	}
	fmt.Printf("Successfully loaded %d templates.\n", len(TemplatePaths))
}

// Define a handler function to render templates with nonce
func renderWithNonce(filePath string) fiber.Handler {
	return func(c fiber.Ctx) error {
		// Retrieve nonce from context
		nonce := c.Locals("nonce").(string)
		fmt.Println("Generated Nonce:", nonce)

		// Render the template and pass the nonce
		return c.Render(filePath, fiber.Map{
			"Nonce": nonce,
		})
	}
}

func NonceAndCSPMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		nonce, err := util.GenerateNonce()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		c.Locals("nonce", nonce)

		// Set the CSP headers
		c.Set("Content-Security-Policy", fmt.Sprintf(
			"default-src 'self'; "+
				"script-src 'self' 'nonce-%[1]s' %[2]s; "+
				"style-src 'self' 'nonce-%[1]s' %[3]s; "+
				"img-src 'self' %[4]s; "+
				"font-src 'self' %[5]s; "+
				"frame-src 'self' %[6]s; "+
				"connect-src 'self' %[7]s",
			nonce, ALLOWED_DOMAINS["script"], ALLOWED_DOMAINS["style"], ALLOWED_DOMAINS["img"], ALLOWED_DOMAINS["font"], ALLOWED_DOMAINS["frame"], ALLOWED_DOMAINS["connect"],
		))
		return c.Next()
	}
}
