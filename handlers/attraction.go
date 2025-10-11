package handlers

import (
	"fmt"
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// ATTRACTION MANAGEMENT - ADMIN
// =============================================================================

// CreateAttraction creates a new attraction
func CreateAttraction(c *fiber.Ctx) error {
	var attraction models.Attraction
	if err := c.BodyParser(&attraction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&attraction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create attraction",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(attraction)
}

// UpdateAttraction updates an existing attraction
func UpdateAttraction(c *fiber.Ctx) error {
	id := c.Params("id")

	var attraction models.Attraction
	if err := config.DB.Where("id = ?", id).First(&attraction).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Attraction not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&attraction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&attraction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update attraction",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(attraction)
}

// DeleteAttraction deletes an attraction
func DeleteAttraction(c *fiber.Ctx) error {
	id := c.Params("id")

	// First, get the attraction to retrieve URLs for R2 deletion
	var attraction models.Attraction
	if err := config.DB.Where("id = ?", id).First(&attraction).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Attraction not found",
			"code":    "NOT_FOUND",
		})
	}

	// Delete from database
	if err := config.DB.Where("id = ?", id).Delete(&models.Attraction{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete attraction",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Delete images from R2 if they are from our R2 bucket
	if utils.Storage != nil {
		// Delete main image
		if err := utils.Storage.DeleteImageIfR2(attraction.ImageURL); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete image from R2: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Attraction deleted successfully",
	})
}

// UpdateGeneralAttractionContent updates the general attraction content (singleton)
func UpdateGeneralAttractionContent(c *fiber.Ctx) error {
	var content models.GeneralAttractionContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.GeneralAttractionContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create general attraction content",
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
			"message": "Failed to update general attraction content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// ATTRACTION MANAGEMENT - PUBLIC
// =============================================================================

// GetAttractions returns all attractions with optional filtering
func GetAttractions(c *fiber.Ctx) error {
	var attractions []models.Attraction
	query := config.DB.Model(&models.Attraction{})

	// Apply featured filter if specified
	if active := c.Query("active"); active == "true" {
		query = query.Where("active = ?", true)
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

	// Count active attractions
	var activeCount int64
	config.DB.Model(&models.Attraction{}).Where("active = ?", true).Count(&activeCount)

	return c.JSON(fiber.Map{
		"data": attractions,
		"meta": fiber.Map{
			"total":        len(attractions),
			"active_count": activeCount,
		},
	})
}

// GetGeneralAttractionContent returns the general attraction content
func GetGeneralAttractionContent(c *fiber.Ctx) error {
	var content models.GeneralAttractionContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.GeneralAttractionContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}
