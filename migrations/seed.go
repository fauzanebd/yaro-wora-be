package migrations

import (
	"encoding/json"
	"log"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"gorm.io/datatypes"
)

// toJSON converts any value to datatypes.JSON
func toJSON(v interface{}) datatypes.JSON {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return datatypes.JSON("{}")
	}
	return datatypes.JSON(jsonBytes)
}

// SeedData creates initial data for the application
func SeedData() {
	db := config.DB

	// Create admin user if not exists
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)
	if userCount == 0 {
		adminUser := models.User{
			Username: config.AppConfig.AdminUsername,
			Password: config.AppConfig.AdminPassword,
			Role:     "super_admin",
			IsActive: true,
		}
		if err := db.Create(&adminUser).Error; err != nil {
			log.Printf("Failed to create admin user: %v", err)
		} else {
			log.Println("✅ Admin user created successfully")
		}
	}

	// Create default pricing if not exists
	var pricingCount int64
	db.Model(&models.Pricing{}).Count(&pricingCount)
	if pricingCount == 0 {
		pricingData := []models.Pricing{
			{
				Type:        "domestic",
				Title:       "Domestic",
				Subtitle:    "Entrance Fee",
				AdultPrice:  60000,
				InfantPrice: 30000,
				Currency:    "IDR",
				Description: "For Indonesian citizens",
			},
			{
				Type:        "locals_sumba",
				Title:       "Locals Sumba",
				Subtitle:    "Entrance Fee",
				AdultPrice:  40000,
				InfantPrice: 20000,
				Currency:    "IDR",
				Description: "For local Sumba residents",
			},
			{
				Type:        "foreigner",
				Title:       "Foreigner",
				Subtitle:    "Entrance Fee",
				AdultPrice:  100000,
				InfantPrice: 50000,
				Currency:    "IDR",
				Description: "For international visitors",
			},
		}

		for _, pricing := range pricingData {
			if err := db.Create(&pricing).Error; err != nil {
				log.Printf("Failed to create pricing: %v", err)
			}
		}
		log.Println("✅ Default pricing data created")
	}

	// Create default gallery categories if not exists
	var galleryCatCount int64
	db.Model(&models.GalleryCategory{}).Count(&galleryCatCount)
	if galleryCatCount == 0 {
		galleryCategories := []models.GalleryCategory{
			{
				ID:            "nature",
				Name:          "Nature",
				NameID:        "Alam",
				Description:   "Natural landscapes and wildlife photography",
				DescriptionID: "Fotografi lanskap alam dan satwa liar",
				Color:         "#22c55e",
				Icon:          "nature",
				SortOrder:     1,
			},
			{
				ID:            "culture",
				Name:          "Culture",
				NameID:        "Budaya",
				Description:   "Cultural ceremonies and traditional activities",
				DescriptionID: "Upacara budaya dan aktivitas tradisional",
				Color:         "#8a0604",
				Icon:          "culture",
				SortOrder:     2,
			},
			{
				ID:            "village",
				Name:          "Village Life",
				NameID:        "Kehidupan Desa",
				Description:   "Daily life and village architecture",
				DescriptionID: "Kehidupan sehari-hari dan arsitektur desa",
				Color:         "#dc2626",
				Icon:          "village",
				SortOrder:     3,
			},
			{
				ID:            "activities",
				Name:          "Activities",
				NameID:        "Aktivitas",
				Description:   "Tourism activities and experiences",
				DescriptionID: "Aktivitas pariwisata dan pengalaman",
				Color:         "#586d12",
				Icon:          "activities",
				SortOrder:     4,
			},
		}

		for _, category := range galleryCategories {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create gallery category: %v", err)
			}
		}
		log.Println("✅ Gallery categories created")
	}

	// Create default news categories if not exists
	var newsCatCount int64
	db.Model(&models.NewsCategory{}).Count(&newsCatCount)
	if newsCatCount == 0 {
		newsCategories := []models.NewsCategory{
			{
				Key:           "culture",
				Name:          "Culture",
				NameID:        "Budaya",
				Description:   "Cultural events, traditions, and heritage news",
				DescriptionID: "Berita acara budaya, tradisi, dan warisan",
				Color:         "#8a0604",
				Icon:          "cultural-heritage",
				SortOrder:     1,
			},
			{
				Key:           "tourism",
				Name:          "Tourism",
				NameID:        "Pariwisata",
				Description:   "Tourism developments and visitor experiences",
				DescriptionID: "Perkembangan pariwisata dan pengalaman pengunjung",
				Color:         "#586d12",
				Icon:          "tourism",
				SortOrder:     2,
			},
			{
				Key:           "events",
				Name:          "Events",
				NameID:        "Acara",
				Description:   "Community events and celebrations",
				DescriptionID: "Acara komunitas dan perayaan",
				Color:         "#dc2626",
				Icon:          "events",
				SortOrder:     3,
			},
			{
				Key:           "environment",
				Name:          "Environment",
				NameID:        "Lingkungan",
				Description:   "Environmental conservation and sustainability",
				DescriptionID: "Konservasi lingkungan dan keberlanjutan",
				Color:         "#16a34a",
				Icon:          "environment",
				SortOrder:     4,
			},
			{
				Key:           "business",
				Name:          "Business",
				NameID:        "Bisnis",
				Description:   "Local business and economic development",
				DescriptionID: "Bisnis lokal dan pengembangan ekonomi",
				Color:         "#0ea5e9",
				Icon:          "business",
				SortOrder:     5,
			},
		}

		for _, category := range newsCategories {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create news category: %v", err)
			}
		}
		log.Println("✅ News categories created")
	}

	// Create default regulation categories if not exists
	var regCatCount int64
	db.Model(&models.RegulationCategory{}).Count(&regCatCount)
	if regCatCount == 0 {
		regulationCategories := []models.RegulationCategory{
			{
				Key:           "general",
				Name:          "General Rules",
				NameID:        "Aturan Umum",
				Description:   "Basic rules and requirements for all visitors",
				DescriptionID: "Aturan dasar dan persyaratan untuk semua pengunjung",
				Icon:          "rules",
				Color:         "#6b7280",
				SortOrder:     1,
			},
			{
				Key:           "cultural",
				Name:          "Cultural Guidelines",
				NameID:        "Panduan Budaya",
				Description:   "Important cultural etiquette and customs",
				DescriptionID: "Etiket budaya penting dan adat istiadat",
				Icon:          "culture",
				Color:         "#8a0604",
				SortOrder:     2,
			},
			{
				Key:           "safety",
				Name:          "Safety & Emergency",
				NameID:        "Keselamatan & Darurat",
				Description:   "Safety protocols and emergency procedures",
				DescriptionID: "Protokol keselamatan dan prosedur darurat",
				Icon:          "shield",
				Color:         "#dc2626",
				SortOrder:     3,
			},
		}

		for _, category := range regulationCategories {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create regulation category: %v", err)
			}
		}
		log.Println("✅ Regulation categories created")
	}

	// Create default contact info if not exists
	var contactCount int64
	db.Model(&models.ContactInfo{}).Count(&contactCount)
	if contactCount == 0 {
		defaultContact := models.ContactInfo{
			Street:     "Yaro Wora Village",
			City:       "East Sumba",
			Province:   "East Nusa Tenggara",
			Country:    "Indonesia",
			PostalCode: "87173",
			Latitude:   -9.6234,
			Longitude:  119.3456,
			Phones:     toJSON([]string{"+62 098 940 974", "+62 903 009 909"}),
			Emails:     toJSON([]string{"info@yarowora.com", "visit@yarowora.com"}),
			WhatsApp:   "+62 903 009 909",
			SocialMedia: toJSON(map[string]string{
				"instagram": "@yarowora_official",
				"facebook":  "Yaro Wora Tourism",
				"youtube":   "Yaro Wora Channel",
			}),
			OperatingHours: toJSON(map[string]string{
				"monday":    "08:00-17:00",
				"tuesday":   "08:00-17:00",
				"wednesday": "08:00-17:00",
				"thursday":  "08:00-17:00",
				"friday":    "08:00-17:00",
				"saturday":  "08:00-16:00",
				"sunday":    "closed",
			}),
		}

		if err := db.Create(&defaultContact).Error; err != nil {
			log.Printf("Failed to create contact info: %v", err)
		} else {
			log.Println("✅ Default contact info created")
		}
	}

	log.Println("🎉 Database seeding completed!")
}
