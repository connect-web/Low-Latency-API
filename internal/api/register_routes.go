package api

import (
	"github.com/connect-web/Low-Latency-API/internal/api/auth"
	"github.com/connect-web/Low-Latency-API/internal/api/middleware"
	v2 "github.com/connect-web/Low-Latency-API/internal/api/v2"
	"github.com/gofiber/fiber/v3"
)

/*
This should contain all of the /api/ routes
with the middleware setup on here.

*/

func CreateRouter(api fiber.Router) {
	// apply Auth rate limits
	authRouter := api.Group("auth", middleware.AuthRateLimits(), auth.CsrfMiddleware)
	auth.CreateRouter(authRouter) // Setup CreateRouter, Login Routes

	// apply regular rate limits
	v2Router := api.Group("v2", middleware.APIRateLimits(), auth.Protected)
	v2.RegisterRouter(v2Router)

}
