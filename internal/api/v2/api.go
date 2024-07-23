package v2

import (
	"github.com/connect-web/Low-Latency-API/internal/api/v2/global"
	"github.com/connect-web/Low-Latency-API/internal/api/v2/ml"
	"github.com/connect-web/Low-Latency-API/internal/api/v2/profile"
	"github.com/connect-web/Low-Latency-API/internal/api/v2/public"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/storage/memory"
	"time"
)

func RegisterRouter(api fiber.Router) {
	cacheStorage := memory.New()

	// Define custom cache middleware
	customCacheMiddleware := cache.New(cache.Config{
		Expiration: 24 * 7 * time.Hour,
		Storage:    cacheStorage,
		Next: func(c fiber.Ctx) bool {
			// Skip caching for the /user/profile route
			return c.Path() == "/api/v2/user/profile" || c.Path() == "/global-stats"
		},
	})
	api.Use(customCacheMiddleware)

	publicRoute := api.Group("/public")

	publicRoute.Get("/skill-toplist", public.GetSkillToplist)
	publicRoute.Get("/skill-toplist-users", public.GetSkillToplistUsers)

	publicRoute.Get("/boss-minigame-toplist", public.GetMinigameToplist)
	publicRoute.Get("/boss-minigame-toplist-users", public.GetMinigameToplistUsers)

	MlRoute := api.Group("/ml")
	MlRoute.Get("/skill-toplist", ml.GetSkillToplist)
	MlRoute.Get("/skill-toplist-users", ml.GetSkillToplistUsers)

	MlRoute.Get("/boss-minigame-toplist", ml.GetMinigameToplist)
	MlRoute.Get("/boss-minigame-toplist-users", ml.GetMinigameToplistUsers)

	userRoute := api.Group("/user")

	userRoute.Get("/profile", profile.GetProfile)
	userRoute.Get("/global-stats", global.GetLowLatencyStats)

}
