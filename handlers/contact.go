package handlers

import (
	"encoding/json"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
)

// =============================================================================
// CONTACT MANAGEMENT - ADMIN
// =============================================================================

// UpdateContactInfo updates the contact info (singleton)
func UpdateContactInfo(c *fiber.Ctx) error {
	var content models.ContactInfo
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.ContactInfo
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create contact info",
				"code":    "INTERNAL_ERROR",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(content)
	}

	// Update existing content
	content.ID = existingContent.ID
	if err := config.DB.Save(&content).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update contact info",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// UpdateContactContent updates the contact content (singleton)
func UpdateContactContent(c *fiber.Ctx) error {
	var content models.ContactContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.ContactContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create general contact content",
				"code":    "INTERNAL_ERROR",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(content)
	}

	// Update existing content
	content.ID = existingContent.ID
	if err := config.DB.Save(&content).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update general contact content",
			"code":    "INTERNAL_ERROR",
		})
	}
	return c.JSON(content)
}

// =============================================================================
// CONTACT MANAGEMENT - PUBLIC
// =============================================================================

// GetContactInfo returns contact information and location details
func GetContactInfo(c *fiber.Ctx) error {
	var contactInfo models.ContactInfo

	if err := config.DB.First(&contactInfo).Error; err != nil {
		jsonBytesPhones, _ := json.Marshal([]string{"+62 098 940 974", "+62 903 009 909"})
		jsonBytesEmails, _ := json.Marshal([]string{"info@yarowora.com", "visit@yarowora.com"})
		jsonBytesSocialMedia, _ := json.Marshal([]models.SocialMedia{
			{Name: "Instagram", Handle: "@yarowora_official", URL: "https://www.instagram.com/yarowora_official/", IconURL: ""},
			{Name: "Facebook", Handle: "Yaro Wora Tourism", URL: "https://www.facebook.com/yarowora.tourism/", IconURL: "https://www.facebook.com/yarowora.tourism/"},
			{Name: "YouTube", Handle: "Yaro Wora Channel", URL: "https://www.youtube.com/channel/UC-9G-_Hw92gR8x6aI6KU4-w", IconURL: "https://www.youtube.com/channel/UC-9G-_Hw92gR8x6aI6KU4-w"},
		})

		// Return default contact info if not found in database
		return c.JSON(fiber.Map{
			"data": models.ContactInfo{
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
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": contactInfo,
	})
}

// GetGeneralContactContent returns the general contact content
func GetGeneralContactContent(c *fiber.Ctx) error {
	var content models.ContactContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.ContactContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}
