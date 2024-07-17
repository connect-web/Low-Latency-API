package v2

import (
	"github.com/connect-web/Low-Latency-API/internal/api/v2/global"
	"github.com/connect-web/Low-Latency-API/internal/api/v2/profile"
	"github.com/connect-web/Low-Latency-API/internal/api/v2/public"
	"github.com/gofiber/fiber/v3"
)

func RegisterRouter(api fiber.Router) {
	publicRoute := api.Group("/public")

	publicRoute.Get("/skill-toplist", public.GetSkillToplist)
	publicRoute.Get("/skill-toplist-users", public.GetSkillToplistUsers)

	publicRoute.Get("/boss-minigame-toplist", public.GetMinigameToplist)
	publicRoute.Get("/boss-minigame-toplist-users", public.GetMinigameToplistUsers)

	userRoute := api.Group("/user")

	userRoute.Get("/profile", profile.GetProfile)
	userRoute.Get("/ban-count", global.GetBanCount)

}
