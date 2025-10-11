package migrations

import (
	"encoding/json"
	"log"
	"time"
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
				Name:          "Nature",
				NameID:        "Alam",
				Description:   "Natural landscapes and wildlife photography",
				DescriptionID: "Fotografi lanskap alam dan satwa liar",
			},
			{
				Name:          "Culture",
				NameID:        "Budaya",
				Description:   "Cultural ceremonies and traditional activities",
				DescriptionID: "Upacara budaya dan aktivitas tradisional",
			},
			{

				Name:          "Village Life",
				NameID:        "Kehidupan Desa",
				Description:   "Daily life and village architecture",
				DescriptionID: "Kehidupan sehari-hari dan arsitektur desa",
			},
			{

				Name:          "Activities",
				NameID:        "Aktivitas",
				Description:   "Tourism activities and experiences",
				DescriptionID: "Aktivitas pariwisata dan pengalaman",
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
				Name:          "Culture",
				NameID:        "Budaya",
				Description:   "Cultural events, traditions, and heritage news",
				DescriptionID: "Berita acara budaya, tradisi, dan warisan",
			},
			{
				Name:          "Tourism",
				NameID:        "Pariwisata",
				Description:   "Tourism developments and visitor experiences",
				DescriptionID: "Perkembangan pariwisata dan pengalaman pengunjung",
			},
			{
				Name:          "Events",
				NameID:        "Acara",
				Description:   "Community events and celebrations",
				DescriptionID: "Acara komunitas dan perayaan",
			},
			{
				Name:          "Environment",
				NameID:        "Lingkungan",
				Description:   "Environmental conservation and sustainability",
				DescriptionID: "Konservasi lingkungan dan keberlanjutan",
			},
			{
				Name:          "Business",
				NameID:        "Bisnis",
				Description:   "Local business and economic development",
				DescriptionID: "Bisnis lokal dan pengembangan ekonomi",
			},
		}

		for _, category := range newsCategories {
			if err := db.Create(&category).Error; err != nil {
				log.Printf("Failed to create news category: %v", err)
			}
		}
		log.Println("✅ News categories created")
	}

	// Create default news authors if not exists
	var newsAuthorCount int64
	db.Model(&models.NewsAuthor{}).Count(&newsAuthorCount)
	if newsAuthorCount == 0 {
		newsAuthors := []models.NewsAuthor{
			{
				Name:   "Yaro Wora Editorial Team",
				Avatar: "https://yarowora.com/images/team/editorial.jpg",
			},
			{
				Name:   "Maria Sumba",
				Avatar: "https://yarowora.com/images/team/maria.jpg",
			},
			{
				Name:   "John Tourism",
				Avatar: "https://yarowora.com/images/team/john.jpg",
			},
		}

		for _, author := range newsAuthors {
			if err := db.Create(&author).Error; err != nil {
				log.Printf("Failed to create news author: %v", err)
			}
		}
		log.Println("✅ News authors created")
	}

	// Create default news page content if not exists
	var newsPageCount int64
	db.Model(&models.NewsPageContent{}).Count(&newsPageCount)
	if newsPageCount == 0 {
		newsPageContent := models.NewsPageContent{
			HeroImageURL:            "https://yarowora.com/images/news/hero.jpg",
			HeroImageThumbnailURL:   "https://yarowora.com/images/news/hero-thumb.jpg",
			Title:                   "Latest News & Updates",
			TitleID:                 "Berita & Update Terbaru",
			Subtitle:                "Stay informed about Yaro Wora Village",
			SubtitleID:              "Tetap terinformasi tentang Desa Yaro Wora",
			HighlightSectionTitle:   "Featured Stories",
			HighlightSectionTitleID: "Cerita Unggulan",
		}

		if err := db.Create(&newsPageContent).Error; err != nil {
			log.Printf("Failed to create news page content: %v", err)
		} else {
			log.Println("✅ News page content created")
		}
	}

	// Create sample news articles if not exists
	var newsArticleCount int64
	db.Model(&models.NewsArticle{}).Count(&newsArticleCount)
	if newsArticleCount == 0 {
		// Get first author and category for sample articles
		var firstAuthor models.NewsAuthor
		var firstCategory models.NewsCategory
		db.First(&firstAuthor)
		db.First(&firstCategory)

		sampleTags := []string{"tourism", "culture", "sumba", "traditional"}
		sampleTagsJSON := toJSON(sampleTags)

		newsArticles := []models.NewsArticle{
			{
				Title:         "Welcome to Yaro Wora Village",
				TitleID:       "Selamat Datang di Desa Yaro Wora",
				Excerpt:       "Discover the beauty and culture of Yaro Wora Village in East Sumba.",
				ExcerptID:     "Temukan keindahan dan budaya Desa Yaro Wora di Sumba Timur.",
				Content:       "# Welcome to Yaro Wora Village\n\nYaro Wora Village is a traditional Sumbanese village located in East Sumba, Indonesia. This village offers visitors a unique opportunity to experience authentic Sumbanese culture, traditional architecture, and warm hospitality.\n\n## What to Expect\n\n- Traditional Sumbanese houses\n- Cultural ceremonies and performances\n- Local handicrafts and textiles\n- Beautiful natural surroundings",
				ContentID:     "# Selamat Datang di Desa Yaro Wora\n\nDesa Yaro Wora adalah desa tradisional Sumba yang terletak di Sumba Timur, Indonesia. Desa ini menawarkan pengunjung kesempatan unik untuk mengalami budaya Sumba yang autentik, arsitektur tradisional, dan keramahan yang hangat.\n\n## Yang Dapat Diharapkan\n\n- Rumah tradisional Sumba\n- Upacara budaya dan pertunjukan\n- Kerajinan tangan dan tekstil lokal\n- Lingkungan alam yang indah",
				AuthorID:      firstAuthor.ID,
				CategoryID:    firstCategory.ID,
				DatePublished: time.Now().AddDate(0, 0, -1), // 1 day ago
				ImageURL:      "https://yarowora.com/images/news/welcome.jpg",
				Tags:          sampleTagsJSON,
				ReadTime:      5,
				IsHeadline:    true,
			},
			{
				Title:         "Traditional Sumbanese Architecture",
				TitleID:       "Arsitektur Tradisional Sumba",
				Excerpt:       "Learn about the unique architectural features of Sumbanese traditional houses.",
				ExcerptID:     "Pelajari tentang fitur arsitektur unik dari rumah tradisional Sumba.",
				Content:       "# Traditional Sumbanese Architecture\n\nThe traditional houses in Yaro Wora Village showcase the unique architectural heritage of Sumba. These houses are built using traditional methods and materials, reflecting the island's rich cultural history.\n\n## Key Features\n\n- High-pitched roofs\n- Intricate wood carvings\n- Traditional building materials\n- Cultural significance",
				ContentID:     "# Arsitektur Tradisional Sumba\n\nRumah tradisional di Desa Yaro Wora menampilkan warisan arsitektur unik Sumba. Rumah-rumah ini dibangun menggunakan metode dan bahan tradisional, mencerminkan sejarah budaya pulau yang kaya.\n\n## Fitur Utama\n\n- Atap tinggi\n- Ukiran kayu yang rumit\n- Bahan bangunan tradisional\n- Signifikansi budaya",
				AuthorID:      firstAuthor.ID,
				CategoryID:    firstCategory.ID,
				DatePublished: time.Now().AddDate(0, 0, -3), // 3 days ago
				ImageURL:      "https://yarowora.com/images/news/architecture.jpg",
				Tags:          sampleTagsJSON,
				ReadTime:      7,
				IsHeadline:    false,
			},
		}

		for _, article := range newsArticles {
			if err := db.Create(&article).Error; err != nil {
				log.Printf("Failed to create news article: %v", err)
			}
		}
		log.Println("✅ Sample news articles created")
	}

	// Create default regulation categories if not exists
	var regCatCount int64
	db.Model(&models.RegulationCategory{}).Count(&regCatCount)
	if regCatCount == 0 {
		regulationCategories := []models.RegulationCategory{
			{

				Name:          "General Rules",
				NameID:        "Aturan Umum",
				Description:   "Basic rules and requirements for all visitors",
				DescriptionID: "Aturan dasar dan persyaratan untuk semua pengunjung",
			},
			{
				Name:          "Cultural Guidelines",
				NameID:        "Panduan Budaya",
				Description:   "Important cultural etiquette and customs",
				DescriptionID: "Etiket budaya penting dan adat istiadat",
			},
			{
				Name:          "Safety & Emergency",
				NameID:        "Keselamatan & Darurat",
				Description:   "Safety protocols and emergency procedures",
				DescriptionID: "Protokol keselamatan dan prosedur darurat",
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
	jsonBytesPhones, _ := json.Marshal([]string{"+62 098 940 974", "+62 903 009 909"})
	jsonBytesEmails, _ := json.Marshal([]string{"info@yarowora.com", "visit@yarowora.com"})
	jsonBytesSocialMedia, _ := json.Marshal([]models.SocialMedia{
		{Name: "Instagram", Handle: "@yarowora_official", URL: "https://www.instagram.com/yarowora_official/", IconURL: ""},
		{Name: "Facebook", Handle: "Yaro Wora Tourism", URL: "https://www.facebook.com/yarowora.tourism/", IconURL: "https://www.facebook.com/yarowora.tourism/"},
		{Name: "YouTube", Handle: "Yaro Wora Channel", URL: "https://www.youtube.com/channel/UC-9G-_Hw92gR8x6aI6KU4-w", IconURL: "https://www.youtube.com/channel/UC-9G-_Hw92gR8x6aI6KU4-w"},
	})
	if contactCount == 0 {
		defaultContact := models.ContactInfo{
			AddressPart1:     "Yaro Wora Village",
			AddressPart1ID:   "Desa Yaro Wora",
			AddressPart2:     "East Sumba, NTT, Indonesia",
			AddressPart2ID:   "Sumba Timur, NTT, Indonesia",
			Latitude:         -9.6234,
			Longitude:        119.3456,
			Phones:           datatypes.JSON(jsonBytesPhones),
			Emails:           datatypes.JSON(jsonBytesEmails),
			SocialMedia:      datatypes.JSON(jsonBytesSocialMedia),
			PlanYourVisitURL: "https://yarowora.com/plan-your-visit",
		}
		if err := db.Create(&defaultContact).Error; err != nil {
			log.Printf("Failed to create contact info: %v", err)
		} else {
			log.Println("✅ Default contact info created")
		}
	}

	log.Println("🎉 Database seeding completed!")
}
