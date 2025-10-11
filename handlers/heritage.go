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
// HERITAGE MANAGEMENT - ADMIN
// =============================================================================

// CreateHeritage creates a new heritage
func CreateHeritage(c *fiber.Ctx) error {
	var heritage models.Heritage
	if err := c.BodyParser(&heritage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&heritage).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create heritage",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(heritage)
}

// UpdateHeritage updates an existing heritage
func UpdateHeritage(c *fiber.Ctx) error {
	id := c.Params("id")

	var heritage models.Heritage
	if err := config.DB.Where("id = ?", id).First(&heritage).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Heritage not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&heritage); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&heritage).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update heritage",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(heritage)
}

// DeleteHeritage deletes a heritage
func DeleteHeritage(c *fiber.Ctx) error {
	id := c.Params("id")

	// First, get the heritage to retrieve URLs for R2 deletion
	var heritage models.Heritage
	if err := config.DB.Where("id = ?", id).First(&heritage).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Heritage not found",
			"code":    "NOT_FOUND",
		})
	}

	// Delete from database
	if err := config.DB.Delete(&models.Heritage{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete heritage",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Delete images from R2 if they are from our R2 bucket
	if utils.Storage != nil {
		// Delete main image and thumbnail
		if err := utils.Storage.DeleteImageWithThumbnailIfR2(heritage.ImageURL); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete image from R2: %v\n", err)
		}

		// Delete images from detail sections
		if err := utils.Storage.DeleteImagesFromDetailSections(string(heritage.HeritageDetailSections)); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete images from heritage detail sections: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Heritage deleted successfully",
	})
}

// UpdateHeritagePageContent updates the heritage page content (singleton)
func UpdateHeritagePageContent(c *fiber.Ctx) error {
	var content models.HeritagePageContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.HeritagePageContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create heritage page content",
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
			"message": "Failed to update heritage page content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// HERITAGE MANAGEMENT - PUBLIC
// =============================================================================

// GetHeritage returns all heritage with filtering options
func GetHeritage(c *fiber.Ctx) error {
	var heritage []models.Heritage
	query := config.DB.Model(&models.Heritage{}).
		Select("id, title, title_id, short_description, short_description_id, image_url, thumbnail_url, sort_order")

	// Apply limit and offset for pagination
	limit := 12
	if l := c.Query("limit"); l != "" {
		if limitInt, err := strconv.Atoi(l); err == nil && limitInt > 0 {
			limit = limitInt
		}
	}

	offset := 0
	if o := c.Query("offset"); o != "" {
		if offsetInt, err := strconv.Atoi(o); err == nil && offsetInt >= 0 {
			offset = offsetInt
		}
	}

	query = query.Limit(limit).Offset(offset)

	if err := query.Order("sort_order ASC, created_at ASC").Find(&heritage).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch heritage data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Convert to summary format
	heritageSummaries := make([]models.HeritageSummary, len(heritage))
	for i, h := range heritage {
		heritageSummaries[i] = models.HeritageSummary{
			ID:                 h.ID,
			Title:              h.Title,
			TitleID:            h.TitleID,
			ShortDescription:   h.ShortDescription,
			ShortDescriptionID: h.ShortDescriptionID,
			ImageURL:           h.ImageURL,
			ThumbnailURL:       h.ThumbnailURL,
			SortOrder:          h.SortOrder,
		}
	}

	// Get total count
	var total int64
	config.DB.Model(&models.Heritage{}).Count(&total)

	// Calculate pagination
	totalPages := 0
	if limit > 0 {
		totalPages = int(total) / limit
		if int(total)%limit != 0 {
			totalPages++
		}
	}
	currentPage := 1
	if limit > 0 {
		currentPage = (offset / limit) + 1
	}

	return c.JSON(fiber.Map{
		"data": heritageSummaries,
		"meta": fiber.Map{
			"total": total,
			"pagination": fiber.Map{
				"current_page": currentPage,
				"per_page":     limit,
				"total_pages":  totalPages,
				"has_next":     currentPage < totalPages,
				"has_previous": currentPage > 1,
			},
		},
	})
}

// GetHeritageByID returns a specific heritage by ID
func GetHeritageByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var heritage models.Heritage
	if err := config.DB.Where("id = ?", id).First(&heritage).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Heritage not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": heritage,
	})
}

// GetHeritagePageContent returns the heritage page content
func GetHeritagePageContent(c *fiber.Ctx) error {
	var content models.HeritagePageContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.HeritagePageContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}
