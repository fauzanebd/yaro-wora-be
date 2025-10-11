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
	api.Get("/profile", handlers.GetProfilePageContent)

	// Destination page content endpoint
	api.Get("/destinations/content", handlers.GetDestinationPageContent)

	// Destinations endpoints
	api.Get("/destinations", handlers.GetDestinations)
	api.Get("/destinations/categories", handlers.GetDestinationCategories)
	api.Get("/destinations/:id", handlers.GetDestinationByID)

	// Gallery page content endpoint
	api.Get("/gallery/content", handlers.GetGalleryPageContent)

	// Gallery endpoints
	api.Get("/gallery", handlers.GetGallery)
	api.Get("/gallery/categories", handlers.GetGalleryCategories)
	api.Get("/gallery/:id", handlers.GetGalleryImageByID)

	// Regulation page content endpoint
	api.Get("/regulations/content", handlers.GetRegulationPageContent)

	// Regulations endpoints
	api.Get("/regulations", handlers.GetRegulations)
	api.Get("/regulations/categories", handlers.GetRegulationCategories)
	api.Get("/regulations/:id", handlers.GetRegulationByID)

	// Facilities page content endpoint
	api.Get("/facilities/content", handlers.GetFacilityPageContent)

	// Facilities endpoints
	api.Get("/facilities", handlers.GetFacilities)
	api.Get("/facilities/categories", handlers.GetFacilityCategories)
	api.Get("/facilities/:id", handlers.GetFacilityByID)

	// News page content endpoint
	api.Get("/news/content", handlers.GetNewsPageContent)

	// News endpoints
	api.Get("/news", handlers.GetNews)
	api.Get("/news/categories", handlers.GetNewsCategories)
	api.Get("/news/authors", handlers.GetNewsAuthors)
	api.Get("/news/authors/:id", handlers.GetNewsAuthorByID)
	api.Get("/news/:id", handlers.GetNewsByID)

	// Heritage page content endpoint
	api.Get("/heritage/content", handlers.GetHeritagePageContent)

	// Heritage endpoints
	api.Get("/heritage", handlers.GetHeritage)
	api.Get("/heritage/:id", handlers.GetHeritageByID)

	// Contact endpoints
	api.Get("/contact-info", handlers.GetContactInfo)
	api.Get("/contact-content", handlers.GetGeneralContactContent)

	// Analytics tracking endpoint (public)
	api.Post("/track", handlers.TrackVisitor)
}
