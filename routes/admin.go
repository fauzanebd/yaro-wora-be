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
	admin.Get("/profile", handlers.GetProfilePageContent)

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
	admin.Put("/profile", handlers.UpdateProfilePageContent)

	// Destination page content management
	admin.Put("/destinations/content", handlers.UpdateDestinationPageContent)

	// Destinations management
	admin.Post("/destinations", handlers.CreateDestination)
	admin.Put("/destinations/:id", handlers.UpdateDestination)
	admin.Delete("/destinations/:id", handlers.DeleteDestination)

	// Destination categories management
	admin.Post("/destinations/categories", handlers.CreateDestinationCategory)
	admin.Put("/destinations/categories/:id", handlers.UpdateDestinationCategory)
	admin.Delete("/destinations/categories/:id", handlers.DeleteDestinationCategory)

	// Gallery page content management
	admin.Put("/gallery/content", handlers.UpdateGalleryPageContent)

	// Destinations management
	admin.Post("/gallery", handlers.CreateGalleryImage)
	admin.Put("/gallery/:id", handlers.UpdateGalleryImage)
	admin.Delete("/gallery/:id", handlers.DeleteGalleryImage)

	// Gallery categories management
	admin.Post("/gallery/categories", handlers.CreateGalleryCategory)
	admin.Put("/gallery/categories/:id", handlers.UpdateGalleryCategory)
	admin.Delete("/gallery/categories/:id", handlers.DeleteGalleryCategory)

	// Regulation page content management
	admin.Put("/regulations/content", handlers.UpdateRegulationPageContent)

	// Regulations management
	admin.Post("/regulations", handlers.CreateRegulation)
	admin.Put("/regulations/:id", handlers.UpdateRegulation)
	admin.Delete("/regulations/:id", handlers.DeleteRegulation)

	// Regulation categories management
	admin.Post("/regulations/categories", handlers.CreateRegulationCategory)
	admin.Put("/regulations/categories/:id", handlers.UpdateRegulationCategory)
	admin.Delete("/regulations/categories/:id", handlers.DeleteRegulationCategory)

	// Facilities page content management
	admin.Put("/facilities/content", handlers.UpdateFacilityPageContent)

	// Facilities management
	admin.Post("/facilities", handlers.CreateFacility)
	admin.Put("/facilities/:id", handlers.UpdateFacility)
	admin.Delete("/facilities/:id", handlers.DeleteFacility)

	// Facility categories management
	admin.Post("/facilities/categories", handlers.CreateFacilityCategory)
	admin.Put("/facilities/categories/:id", handlers.UpdateFacilityCategory)
	admin.Delete("/facilities/categories/:id", handlers.DeleteFacilityCategory)

	// News page content management
	admin.Put("/news/content", handlers.UpdateNewsPageContent)

	// News management
	admin.Post("/news", handlers.CreateNews)
	admin.Put("/news/:id", handlers.UpdateNews)
	admin.Delete("/news/:id", handlers.DeleteNews)

	// News categories management
	admin.Post("/news/categories", handlers.CreateNewsCategory)
	admin.Put("/news/categories/:id", handlers.UpdateNewsCategory)
	admin.Delete("/news/categories/:id", handlers.DeleteNewsCategory)

	// News authors management
	admin.Post("/news/authors", handlers.CreateNewsAuthor)
	admin.Put("/news/authors/:id", handlers.UpdateNewsAuthor)
	admin.Delete("/news/authors/:id", handlers.DeleteNewsAuthor)

	// Contact management
	admin.Put("/contact-info", handlers.UpdateContactInfo)
	admin.Put("/contact-content", handlers.UpdateContactContent)

	// Content management
	admin.Post("/content/upload", handlers.UploadContent)

	// Heritage page content management
	admin.Put("/heritage/content", handlers.UpdateHeritagePageContent)

	// Heritage management
	admin.Post("/heritage", handlers.CreateHeritage)
	admin.Put("/heritage/:id", handlers.UpdateHeritage)
	admin.Delete("/heritage/:id", handlers.DeleteHeritage)

	// Analytics & Reports
	admin.Get("/analytics/storage", handlers.GetStorageAnalytics)
	admin.Get("/analytics/visitors", handlers.GetVisitorAnalytics)
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
