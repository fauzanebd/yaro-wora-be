package routes

import (
	"yaro-wora-be/handlers"
	"yaro-wora-be/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupAdminRoutes sets up all admin routes (protected)
func SetupAdminRoutes(api fiber.Router) {
	// Admin group with authentication middleware
	admin := api.Group("/admin", middleware.AdminAuth())

	// Profile endpoint for authenticated user
	admin.Get("/profile", handlers.Profile)

	// Main page management
	admin.Post("/carousel", handlers.CreateCarousel)
	admin.Put("/carousel/:id", handlers.UpdateCarousel)
	admin.Delete("/carousel/:id", handlers.DeleteCarousel)

	admin.Post("/why-visit", handlers.CreateWhyVisit)
	admin.Put("/why-visit/:id", handlers.UpdateWhyVisit)
	admin.Delete("/why-visit/:id", handlers.DeleteWhyVisit)

	admin.Put("/why-visit-content", handlers.UpdateGeneralWhyVisitContent)

	admin.Post("/selling-points", handlers.CreateSellingPoint)
	admin.Put("/selling-points/:id", handlers.UpdateSellingPoint)
	admin.Delete("/selling-points/:id", handlers.DeleteSellingPoint)

	admin.Post("/attractions", handlers.CreateAttraction)
	admin.Put("/attractions/:id", handlers.UpdateAttraction)
	admin.Delete("/attractions/:id", handlers.DeleteAttraction)

	admin.Put("/attraction-content", handlers.UpdateGeneralAttractionContent)

	admin.Put("/pricing", handlers.UpdatePricing)

	admin.Put("/pricing-content", handlers.UpdateGeneralPricingContent)

	// Profile page management
	admin.Put("/profile", handlers.UpdateProfile)

	// Destinations management
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
}
