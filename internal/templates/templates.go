package templates

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/api/middleware"
	"github.com/connect-web/Low-Latency-API/internal/util"
	"github.com/gofiber/fiber/v3"
	"html/template"
)

func Run(app *fiber.App) {

	app.Get("/log-in-template", templateLogin, middleware.NonceAndCSPMiddleware())
}

func getScripts(nonce string) template.HTML {
	scripts := fmt.Sprintf(`
	<script nonce="%[1]s" src="resources/auth/loading/loading.js"></script>
    <script nonce="%[1]s" src="resources/auth/redirect/redirect.js"></script>
    <script nonce="%[1]s" src="resources/js/unauth-session.js"></script>
    <script nonce="%[1]s" src="resources/auth/auth.js"></script>
    <script nonce="%[1]s" src="resources/auth/login.js"></script>
    <script nonce="%[1]s" src="resources/home/buttons/script.js"></script>
    <script nonce="%[1]s" src="https://js.hcaptcha.com/1/api.js" async defer></script>`, nonce)

	s := template.HTML(scripts)
	fmt.Println(s)
	return s
}

func templateLogin(c fiber.Ctx) error {
	nonce := c.Locals("nonce")
	if nonce == nil {
		fmt.Printf("NONCE WAS NIL...")
		return util.InternalServerError(c)
	}
	nonce = nonce.(string)
	fmt.Printf("Template nonce: %s\n", nonce)
	c.Set("Cache-Control", "no-store")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "0")
	return c.Render("login", fiber.Map{
		"Nonce":   nonce,
		"Scripts": getScripts(nonce.(string)),
	})
}
