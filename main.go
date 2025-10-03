package main

import (
	"log"
	"time"
	"yaro-wora-be/config"
	"yaro-wora-be/handlers"
	"yaro-wora-be/middleware"
	"yaro-wora-be/migrations"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Connect to database
	config.ConnectDatabase()

	// Run migrations
	models.AutoMigrate()

	// Seed initial data
	if config.AppConfig.AppEnv == "development" {
		// Import migrations package
		migrations.SeedData()
	}

	// Initialize storage
	if err := utils.InitStorage(); err != nil {
		log.Printf("Warning: Failed to initialize storage: %v", err)
		log.Println("Some upload features may not work properly")
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error":     true,
				"message":   err.Error(),
				"code":      "FIBER_ERROR",
				"timestamp": time.Now().Format(time.RFC3339),
			})
		},
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"message": "Yaro Wora API is running",
			"version": config.AppConfig.APIVersion,
		})
	})

	// API routes
	api := app.Group("/v1")

	// =============================================================================
	// PUBLIC ROUTES
	// =============================================================================

	// Main page endpoints
	api.Get("/carousel", handlers.GetCarousel)
	api.Get("/attractions", handlers.GetAttractions)
	api.Get("/pricing", handlers.GetPricing)

	// Profile page endpoints
	api.Get("/profile", handlers.GetProfile)

	// Destinations endpoints
	api.Get("/destinations", handlers.GetDestinations)
	api.Get("/destinations/main", handlers.GetMainDestination)
	api.Get("/destinations/:id", handlers.GetDestinationByID)

	// Gallery endpoints
	api.Get("/gallery", handlers.GetGallery)
	api.Get("/gallery/categories", handlers.GetGalleryCategories)
	api.Get("/gallery/:id", handlers.GetGalleryImageByID)

	// Regulations endpoints
	api.Get("/regulations", handlers.GetRegulations)
	api.Get("/regulations/categories", handlers.GetRegulationCategories)
	api.Get("/regulations/:id", handlers.GetRegulationByID)

	// Facilities endpoints
	api.Get("/facilities", handlers.GetFacilities)
	api.Get("/facilities/:id", handlers.GetFacilityByID)
	api.Post("/facilities/:id/book", handlers.BookFacility)

	// News endpoints
	api.Get("/news", handlers.GetNews)
	api.Get("/news/categories", handlers.GetNewsCategories)
	api.Get("/news/featured", handlers.GetFeaturedNews)
	api.Get("/news/:id", handlers.GetNewsByID)

	// Contact endpoints
	api.Post("/contact", handlers.SubmitContact)
	api.Get("/contact-info", handlers.GetContactInfo)

	// =============================================================================
	// AUTHENTICATION ROUTES
	// =============================================================================

	// Simple login for basic auth tier
	api.Post("/auth/login", handlers.SimpleLogin)

	// JWT-based login for advanced auth (if needed)
	api.Post("/auth/jwt-login", handlers.Login)

	// =============================================================================
	// ADMIN ROUTES (Protected)
	// =============================================================================

	// Admin group with authentication middleware
	admin := api.Group("/admin", middleware.AdminAuth())

	// Profile endpoint for authenticated user
	admin.Get("/profile", handlers.Profile)

	// Main page management
	admin.Post("/carousel", handlers.CreateCarousel)
	admin.Put("/carousel/:id", handlers.UpdateCarousel)
	admin.Delete("/carousel/:id", handlers.DeleteCarousel)

	admin.Post("/attractions", handlers.CreateAttraction)
	admin.Put("/attractions/:id", handlers.UpdateAttraction)
	admin.Delete("/attractions/:id", handlers.DeleteAttraction)

	admin.Put("/pricing", handlers.UpdatePricing)

	// Profile page management
	admin.Put("/profile", handlers.UpdateProfile)

	// Destinations management (you can implement these similar to attractions)
	admin.Post("/destinations", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Create destination - TODO"})
	})
	admin.Put("/destinations/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update destination - TODO"})
	})
	admin.Delete("/destinations/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete destination - TODO"})
	})

	// Gallery management
	admin.Post("/gallery", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Upload gallery image - TODO"})
	})
	admin.Put("/gallery/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update gallery image - TODO"})
	})
	admin.Delete("/gallery/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete gallery image - TODO"})
	})

	admin.Post("/gallery/categories", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Create gallery category - TODO"})
	})
	admin.Put("/gallery/categories/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update gallery category - TODO"})
	})
	admin.Delete("/gallery/categories/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete gallery category - TODO"})
	})

	// Regulations management
	admin.Post("/regulations", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Create regulation - TODO"})
	})
	admin.Put("/regulations/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update regulation - TODO"})
	})
	admin.Delete("/regulations/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete regulation - TODO"})
	})

	// Facilities management
	admin.Post("/facilities", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Create facility - TODO"})
	})
	admin.Put("/facilities/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update facility - TODO"})
	})
	admin.Delete("/facilities/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete facility - TODO"})
	})
	admin.Get("/facilities/:id/bookings", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Get facility bookings - TODO"})
	})

	// News management
	admin.Post("/news", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Create news article - TODO"})
	})
	admin.Put("/news/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update news article - TODO"})
	})
	admin.Delete("/news/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete news article - TODO"})
	})

	// Contact & Booking management
	admin.Get("/contacts", handlers.GetContacts)
	admin.Get("/contacts/:id", handlers.GetContactByID)
	admin.Put("/contacts/:id", handlers.UpdateContactStatus)

	admin.Get("/bookings", handlers.GetBookings)
	admin.Put("/bookings/:id", handlers.UpdateBookingStatus)

	// Content management
	admin.Put("/contact-info", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update contact info - TODO"})
	})
	admin.Post("/content/upload", handlers.UploadContent)

	// Analytics & Reports (placeholder)
	admin.Get("/analytics/visitors", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Visitor analytics - TODO"})
	})
	admin.Get("/analytics/bookings", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Booking analytics - TODO"})
	})
	admin.Get("/analytics/content", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Content analytics - TODO"})
	})

	// User management (placeholder)
	admin.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "List users - TODO"})
	})
	admin.Post("/users", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Create user - TODO"})
	})
	admin.Put("/users/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Update user - TODO"})
	})
	admin.Delete("/users/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Delete user - TODO"})
	})

	// =============================================================================
	// SIMPLE AUTH ROUTES (Alternative for basic tier)
	// =============================================================================

	// Simple admin routes with basic auth
	simpleAdmin := api.Group("/simple-admin", middleware.BasicAuth())

	simpleAdmin.Get("/dashboard", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to simple admin dashboard",
			"user":    c.Locals("basicAuthUser"),
		})
	})

	// Start server
	port := ":" + config.AppConfig.Port
	log.Printf("🚀 Server starting on port %s", config.AppConfig.Port)
	log.Printf("📚 API Documentation available at: http://localhost%s/health", port)
	log.Printf("🔧 Environment: %s", config.AppConfig.AppEnv)

	if err := app.Listen(port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
