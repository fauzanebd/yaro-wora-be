package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes initializes all routes for the application
func SetupRoutes(app *fiber.App) {
	// API routes
	api := app.Group("/v1")

	// Setup public routes
	SetupPublicRoutes(api)

	// Setup authentication routes
	SetupAuthRoutes(api)

	// Setup admin routes
	SetupAdminRoutes(api)
}
