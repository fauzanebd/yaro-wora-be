package routes

import (
	"yaro-wora-be/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupPublicRoutes sets up all public routes
func SetupPublicRoutes(api fiber.Router) {
	// Main page endpoints
	api.Get("/carousel", handlers.GetCarousel)
	api.Get("/why-visit", handlers.GetWhyVisit)
	api.Get("/why-visit-content", handlers.GetGeneralWhyVisitContent)
	api.Get("/selling-points", handlers.GetSellingPoints)
	api.Get("/attractions", handlers.GetAttractions)
	api.Get("/attraction-content", handlers.GetGeneralAttractionContent)
	api.Get("/pricing", handlers.GetPricing)
	api.Get("/pricing-content", handlers.GetGeneralPricingContent)

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
}
