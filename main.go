package main

import (
	"log"
	"time"
	"yaro-wora-be/config"
	"yaro-wora-be/migrations"
	"yaro-wora-be/models"
	"yaro-wora-be/routes"
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
	// if config.AppConfig.AppEnv == "development" {
	// Import migrations package
	migrations.SeedData()
	// }

	// Initialize storage
	if err := utils.InitStorage(); err != nil {
		log.Printf("Warning: Failed to initialize storage: %v", err)
		log.Println("Some upload features may not work properly")
	}

	// Create Fiber app
	log.Printf("📦 Max file upload size configured: %d bytes (%.2f MB)",
		config.AppConfig.MaxFileUploadSize,
		float64(config.AppConfig.MaxFileUploadSize)/(1024*1024))

	app := fiber.New(fiber.Config{
		BodyLimit: config.AppConfig.MaxFileUploadSize, // Set max body size from config (default 20MB)
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
		AllowHeaders: "Origin,Content-Type,Accept,Authorization,X-Session-ID",
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"message": "Yaro Wora API is running",
			"version": config.AppConfig.APIVersion,
		})
	})

	// Setup all routes
	routes.SetupRoutes(app)

	// Start server
	port := ":" + config.AppConfig.Port
	log.Printf("🚀 Server starting on port %s", config.AppConfig.Port)
	log.Printf("Health check available at: http://localhost%s/health", port)
	log.Printf("🔧 Environment: %s", config.AppConfig.AppEnv)

	if err := app.Listen(port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
