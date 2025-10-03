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
		&Attraction{},
		&Pricing{},
		&Profile{},

		// Destination models
		&Destination{},

		// Gallery models
		&GalleryCategory{},
		&GalleryImage{},

		// Regulation models
		&RegulationCategory{},
		&Regulation{},

		// Facility models
		&Facility{},
		&Booking{},

		// News models
		&NewsCategory{},
		&NewsArticle{},

		// Contact models
		&ContactSubmission{},
		&ContactInfo{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
}
