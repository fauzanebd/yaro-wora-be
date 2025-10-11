package handlers

import (
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// WHY VISIT MANAGEMENT - ADMIN
// =============================================================================

// CreateWhyVisit creates a new why visit item
func CreateWhyVisit(c *fiber.Ctx) error {
	var whyVisit models.WhyVisit
	if err := c.BodyParser(&whyVisit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create why visit item",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(whyVisit)
}

// UpdateWhyVisit updates an existing why visit item
func UpdateWhyVisit(c *fiber.Ctx) error {
	id := c.Params("id")

	var whyVisit models.WhyVisit
	if err := config.DB.Where("id = ?", id).First(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Why visit item not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&whyVisit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update why visit item",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(whyVisit)
}

// DeleteWhyVisit deletes a why visit item
func DeleteWhyVisit(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.WhyVisit{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete why visit item",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Why visit item deleted successfully",
	})
}

// UpdateGeneralWhyVisitContent updates the general why visit content (singleton)
func UpdateGeneralWhyVisitContent(c *fiber.Ctx) error {
	var content models.GeneralWhyVisitContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.GeneralWhyVisitContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create general why visit content",
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
			"message": "Failed to update general why visit content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// WHY VISIT MANAGEMENT - PUBLIC
// =============================================================================

// GetWhyVisit returns all why visit items
func GetWhyVisit(c *fiber.Ctx) error {
	var whyVisit []models.WhyVisit

	if err := config.DB.Order("created_at ASC").Find(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch why visit data",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"data": whyVisit,
		"meta": fiber.Map{
			"total": len(whyVisit),
		},
	})
}

// GetGeneralWhyVisitContent returns the general why visit content
func GetGeneralWhyVisitContent(c *fiber.Ctx) error {
	var content models.GeneralWhyVisitContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.GeneralWhyVisitContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}
