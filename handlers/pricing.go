package handlers

import (
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// PRICING MANAGEMENT - ADMIN
// =============================================================================

// UpdatePricing updates entrance fee pricing
func UpdatePricing(c *fiber.Ctx) error {
	var pricing models.Pricing
	if err := c.BodyParser(&pricing); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing pricing by type, create if not exists
	var existingPricing models.Pricing
	if err := config.DB.Where("type = ?", pricing.Type).First(&existingPricing).Error; err != nil {
		// Create new pricing
		if err := config.DB.Create(&pricing).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create pricing",
				"code":    "INTERNAL_ERROR",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(pricing)
	}

	// Update existing pricing
	pricing.ID = existingPricing.ID
	if err := config.DB.Save(&pricing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update pricing",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(pricing)
}

// UpdateGeneralPricingContent updates the general pricing content (singleton)
func UpdateGeneralPricingContent(c *fiber.Ctx) error {
	var content models.GeneralPricingContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.GeneralPricingContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create general pricing content",
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
			"message": "Failed to update general pricing content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// PRICING MANAGEMENT - PUBLIC
// =============================================================================

// GetPricing returns entrance fee pricing
func GetPricing(c *fiber.Ctx) error {
	var pricings []models.Pricing

	if err := config.DB.Find(&pricings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch pricing data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Transform to the expected format
	results := make([]fiber.Map, len(pricings))
	for i, pricing := range pricings {
		results[i] = fiber.Map{
			"type":                 pricing.Type,
			"title":                pricing.Title,
			"title_id":             pricing.TitleID,
			"subtitle":             pricing.Subtitle,
			"subtitle_id":          pricing.SubtitleID,
			"adult_price":          pricing.AdultPrice,
			"infant_price":         pricing.InfantPrice,
			"currency":             pricing.Currency,
			"description":          pricing.Description,
			"image_url":            pricing.ImageURL,
			"thumbnail_url":        pricing.ThumbnailURL,
			"color":                pricing.PrimaryColor,
			"start_gradient_color": pricing.StartGradientColor,
			"end_gradient_color":   pricing.EndGradientColor,
			"created_at":           pricing.CreatedAt,
			"updated_at":           pricing.UpdatedAt,
		}
	}

	return c.JSON(fiber.Map{
		"data": results,
	})
}

// GetGeneralPricingContent returns the general pricing content
func GetGeneralPricingContent(c *fiber.Ctx) error {
	var content models.GeneralPricingContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.GeneralPricingContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}
