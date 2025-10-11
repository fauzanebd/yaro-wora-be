package models

import (
	"log"
	"yaro-wora-be/config"
)

// AutoMigrate runs database migrations for all models
func AutoMigrate() {
	db := config.DB

	err := db.AutoMigrate(
		// Core models
		&User{},
		&Carousel{},
		&WhyVisit{},
		&GeneralWhyVisitContent{},
		&SellingPoint{},
		&Attraction{},
		&GeneralAttractionContent{},
		&Pricing{},
		&GeneralPricingContent{},

		// Profile Page Content models
		&ProfilePageContent{},

		// Destination models
		&Destination{},
		&DestinationDetailSection{},
		&DestinationPageContent{},
		&DestinationCategory{},

		// Gallery models
		&GalleryPageContent{},
		&GalleryImage{},
		&GalleryCategory{},

		// Regulation models
		&RegulationPageContent{},
		&RegulationCategory{},
		&Regulation{},

		// Facility models
		&Facility{},
		&FacilityCategory{},
		&FacilityPageContent{},
		&FacilityDetailSection{},

		// News models
		&NewsArticle{},
		&NewsCategory{},
		&NewsAuthor{},
		&NewsPageContent{},

		// Heritage models
		&Heritage{},
		&HeritagePageContent{},

		// Contact models
		&ContactContent{},
		&ContactInfo{},

		// Analytics models
		&Visitor{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
}
