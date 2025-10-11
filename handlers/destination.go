package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// DESTINATION MANAGEMENT - ADMIN
// =============================================================================

// CreateDestination creates a new destination
func CreateDestination(c *fiber.Ctx) error {
	var destination models.Destination
	if err := c.BodyParser(&destination); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&destination).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create destination",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(destination)
}

// UpdateDestination updates an existing destination
func UpdateDestination(c *fiber.Ctx) error {
	id := c.Params("id")

	var destination models.Destination
	if err := config.DB.Where("id = ?", id).First(&destination).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Destination not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&destination); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&destination).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update destination",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(destination)
}

// DeleteDestination deletes a destination
func DeleteDestination(c *fiber.Ctx) error {
	id := c.Params("id")

	// First, get the destination to retrieve URLs for R2 deletion
	var destination models.Destination
	if err := config.DB.Where("id = ?", id).First(&destination).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Destination not found",
			"code":    "NOT_FOUND",
		})
	}

	// Delete from database
	if err := config.DB.Delete(&models.Destination{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete destination",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Delete images from R2 if they are from our R2 bucket
	if utils.Storage != nil {
		// Delete main image
		if err := utils.Storage.DeleteImageWithThumbnailIfR2(destination.ImageURL); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete image from R2: %v\n", err)
		}

		// Delete images from detail sections
		if err := utils.Storage.DeleteImagesFromDetailSections(string(destination.DestinationDetailSections)); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete images from destination detail sections: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Destination deleted successfully",
	})
}

// UpdateDestinationPageContent updates the destination page content (singleton)
func UpdateDestinationPageContent(c *fiber.Ctx) error {
	var content models.DestinationPageContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.DestinationPageContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create destination page content",
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
			"message": "Failed to update destination page content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// DESTINATION CATEGORY MANAGEMENT - ADMIN
// =============================================================================

// CreateDestinationCategory creates a new destination category
func CreateDestinationCategory(c *fiber.Ctx) error {
	var category models.DestinationCategory
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create destination category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// UpdateDestinationCategory updates an existing destination category
func UpdateDestinationCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.DestinationCategory
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Destination category not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update destination category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(category)
}

// DeleteDestinationCategory deletes a destination category
func DeleteDestinationCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.DestinationCategory{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete destination category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Destination category deleted successfully",
	})
}

// =============================================================================
// DESTINATION MANAGEMENT - PUBLIC
// =============================================================================

// GetDestinations returns all destinations with filtering options
func GetDestinations(c *fiber.Ctx) error {
	var destinations []models.Destination
	query := config.DB.Model(&models.Destination{}).
		Select("id, title, title_id, short_description, short_description_id, image_url, thumbnail_url, highlights, highlights_id, is_featured, sort_order, category_id").
		Preload("DestinationCategory")

	if isFeatured := c.Query("featured"); isFeatured == "true" {
		query = query.Where("is_featured = ?", true)
	}

	// Apply category filter
	// Support comma-separated list, e.g. category=1,2,3
	if category := c.Query("category"); category != "" {
		cats := make([]string, 0)
		for _, part := range strings.Split(category, ",") {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				cats = append(cats, trimmed)
			}
		}
		if len(cats) == 1 {
			query = query.Where("category_id = ?", cats[0])
		} else if len(cats) > 1 {
			query = query.Where("category_id IN ?", cats)
		}
	}

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

	if err := query.Order("sort_order ASC, created_at ASC, is_featured DESC").Find(&destinations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch destinations data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Convert to summary format
	destinationSummaries := make([]models.DestinationSummary, len(destinations))
	for i, dest := range destinations {
		destinationSummaries[i] = models.DestinationSummary{
			ID:                  dest.ID,
			Title:               dest.Title,
			TitleID:             dest.TitleID,
			ShortDescription:    dest.ShortDescription,
			ShortDescriptionID:  dest.ShortDescriptionID,
			ImageURL:            dest.ImageURL,
			ThumbnailURL:        dest.ThumbnailURL,
			Highlights:          dest.Highlights,
			HighlightsID:        dest.HighlightsID,
			IsFeatured:          dest.IsFeatured,
			SortOrder:           dest.SortOrder,
			CategoryID:          dest.CategoryID,
			DestinationCategory: dest.DestinationCategory,
		}
	}

	// Get total count (apply same filters)
	var total int64
	countQuery := config.DB.Model(&models.Destination{})
	if isFeatured := c.Query("featured"); isFeatured == "true" {
		countQuery = countQuery.Where("is_featured = ?", true)
	}
	if category := c.Query("category"); category != "" {
		cats := make([]string, 0)
		for _, part := range strings.Split(category, ",") {
			trimmed := strings.TrimSpace(part)
			if trimmed != "" {
				cats = append(cats, trimmed)
			}
		}
		if len(cats) == 1 {
			countQuery = countQuery.Where("category_id = ?", cats[0])
		} else if len(cats) > 1 {
			countQuery = countQuery.Where("category_id IN ?", cats)
		}
	}
	countQuery.Count(&total)

	// Get categories with counts from DestinationCategory table
	categories := getDestinationCategoriesFromTable()

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
		"data": destinationSummaries,
		"meta": fiber.Map{
			"total":      total,
			"categories": categories,
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

// GetDestinationByID returns a specific destination by ID
func GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var destination models.Destination
	if err := config.DB.Preload("DestinationCategory").Where("id = ?", id).First(&destination).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Destination not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": destination,
	})
}

// GetDestinationCategories returns all destination categories with counts
func GetDestinationCategories(c *fiber.Ctx) error {
	categories := getDestinationCategoriesFromTable()

	var totalDestinations int64
	config.DB.Model(&models.Destination{}).Count(&totalDestinations)

	return c.JSON(fiber.Map{
		"data": categories,
		"meta": fiber.Map{
			"total_categories":   len(categories),
			"total_destinations": totalDestinations,
		},
	})
}

// GetDestinationPageContent returns the destination page content
func GetDestinationPageContent(c *fiber.Ctx) error {
	var content models.DestinationPageContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.DestinationPageContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}

// getDestinationCategoriesFromTable returns categories from DestinationCategory table with destination counts
func getDestinationCategoriesFromTable() []fiber.Map {
	var catRows []models.DestinationCategory
	if err := config.DB.Find(&catRows).Error; err != nil {
		return []fiber.Map{}
	}

	categories := make([]fiber.Map, 0, len(catRows))
	for _, cat := range catRows {
		var count int64
		config.DB.Model(&models.Destination{}).Where("category_id = ? AND is_featured = ?", cat.ID, false).Count(&count)

		categories = append(categories, fiber.Map{
			"id":             cat.ID,
			"name":           cat.Name,
			"name_id":        cat.NameID,
			"description":    cat.Description,
			"description_id": cat.DescriptionID,
			"count":          count,
		})
	}
	return categories
}
