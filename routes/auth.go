package routes

import (
	"yaro-wora-be/handlers"
	"yaro-wora-be/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes sets up all authentication routes
func SetupAuthRoutes(api fiber.Router) {

	// JWT-based login for advanced auth (if needed)
	api.Post("/auth/login", handlers.Login)

	// Simple admin routes with basic auth
	simpleAdmin := api.Group("/simple-admin", middleware.BasicAuth())

	simpleAdmin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to simple admin dashboard",
			"user":    c.Locals("basicAuthUser"),
		})
	})
}
