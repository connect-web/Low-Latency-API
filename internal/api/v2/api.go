package v2

import (
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/api/v2/public"
	"github.com/gofiber/fiber/v3"
)

func RegisterRouter(api fiber.Router) {
	v2 := api.Group("/v2", auth.Protected)

	publicRoute := v2.Group("/public")

	publicRoute.Get("/skill-toplist", public.GetSkillToplist)
	publicRoute.Get("/skill-toplist-users", public.GetSkillToplistUsers)

	publicRoute.Get("/boss-minigame-toplist", public.GetMinigameToplist)
	publicRoute.Get("/boss-minigame-toplist-users", public.GetMinigameToplistUsers)

}
