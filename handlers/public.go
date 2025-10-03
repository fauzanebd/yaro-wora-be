package handlers

import (
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// GetCarousel returns all active carousel slides
func GetCarousel(c *fiber.Ctx) error {
	var carousels []models.Carousel

	if err := config.DB.Where("is_active = ?", true).Order("order ASC, created_at ASC").Find(&carousels).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch carousel data",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"data": carousels,
		"meta": fiber.Map{
			"total":              len(carousels),
			"auto_play_interval": 6000,
		},
	})
}

// GetAttractions returns all attractions with optional filtering
func GetAttractions(c *fiber.Ctx) error {
	var attractions []models.Attraction
	query := config.DB.Model(&models.Attraction{})

	// Apply featured filter if specified
	if featured := c.Query("featured"); featured == "true" {
		query = query.Where("is_featured = ?", true)
	}

	// Apply limit
	if limit := c.Query("limit"); limit != "" {
		if limitInt, err := strconv.Atoi(limit); err == nil && limitInt > 0 {
			query = query.Limit(limitInt)
		}
	}

	if err := query.Order("sort_order ASC, created_at ASC").Find(&attractions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch attractions data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Count featured attractions
	var featuredCount int64
	config.DB.Model(&models.Attraction{}).Where("is_featured = ?", true).Count(&featuredCount)

	return c.JSON(fiber.Map{
		"data": attractions,
		"meta": fiber.Map{
			"total":          len(attractions),
			"featured_count": featuredCount,
		},
	})
}

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
	data := make(map[string]interface{})
	for _, pricing := range pricings {
		data[pricing.Type] = fiber.Map{
			"title":        pricing.Title,
			"subtitle":     pricing.Subtitle,
			"adult_price":  pricing.AdultPrice,
			"infant_price": pricing.InfantPrice,
			"currency":     pricing.Currency,
			"description":  pricing.Description,
		}
	}

	return c.JSON(fiber.Map{
		"data":         data,
		"last_updated": "2024-01-15T10:30:00Z", // You might want to track this properly
	})
}

// GetProfile returns village profile information
func GetProfile(c *fiber.Ctx) error {
	var profile models.Profile

	if err := config.DB.First(&profile).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Profile not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"title":       profile.Title,
			"description": profile.Description,
			"vision": fiber.Map{
				"title":   profile.VisionTitle,
				"content": profile.VisionContent,
			},
			"mission": fiber.Map{
				"title":   profile.MissionTitle,
				"content": profile.MissionContent,
			},
			"objectives": fiber.Map{
				"title":   profile.ObjectivesTitle,
				"content": profile.ObjectivesContent,
			},
			"featured_images": profile.FeaturedImages,
			"last_updated":    profile.UpdatedAt,
		},
	})
}
