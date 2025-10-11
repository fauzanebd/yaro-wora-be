package handlers

import (
	"strconv"
	"strings"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// REGULATION MANAGEMENT - ADMIN
// =============================================================================

// CreateRegulation creates a new regulation
func CreateRegulation(c *fiber.Ctx) error {
	var regulation models.Regulation
	if err := c.BodyParser(&regulation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&regulation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create regulation",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(regulation)
}

// UpdateRegulation updates an existing regulation
func UpdateRegulation(c *fiber.Ctx) error {
	id := c.Params("id")

	var regulation models.Regulation
	if err := config.DB.Where("id = ?", id).First(&regulation).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Regulation not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&regulation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&regulation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update regulation",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(regulation)
}

// DeleteRegulation deletes a regulation
func DeleteRegulation(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.Regulation{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete regulation",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Regulation deleted successfully",
	})
}

// UpdateRegulationPageContent updates the regulation page content (singleton)
func UpdateRegulationPageContent(c *fiber.Ctx) error {
	var content models.RegulationPageContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.RegulationPageContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create regulation page content",
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
			"message": "Failed to update regulation page content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// REGULATION CATEGORY MANAGEMENT - ADMIN
// =============================================================================

// CreateRegulationCategory creates a new regulation category
func CreateRegulationCategory(c *fiber.Ctx) error {
	var category models.RegulationCategory
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
			"message": "Failed to create regulation category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// UpdateRegulationCategory updates an existing regulation category
func UpdateRegulationCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.RegulationCategory
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Regulation category not found",
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
			"message": "Failed to update regulation category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(category)
}

// DeleteRegulationCategory deletes a regulation category
func DeleteRegulationCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.RegulationCategory{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete regulation category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Regulation category deleted successfully",
	})
}

// =============================================================================
// REGULATION MANAGEMENT - PUBLIC
// =============================================================================

// GetRegulations returns all regulations with filtering options
func GetRegulations(c *fiber.Ctx) error {
	var regulations []models.Regulation
	query := config.DB.Model(&models.Regulation{}).
		Preload("RegulationCategory")

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

	if err := query.Order("created_at ASC").Find(&regulations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch regulations data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Get total count (apply same filters)
	var total int64
	countQuery := config.DB.Model(&models.Regulation{})
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

	// Get categories with counts from RegulationCategory table
	categories := getRegulationCategoriesFromTable()

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
		"data": regulations,
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

// GetRegulationByID returns a specific regulation by ID
func GetRegulationByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var regulation models.Regulation
	if err := config.DB.Preload("RegulationCategory").Where("id = ?", id).First(&regulation).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Regulation not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": regulation,
	})
}

// GetRegulationCategories returns all regulation categories with counts
func GetRegulationCategories(c *fiber.Ctx) error {
	categories := getRegulationCategoriesFromTable()

	var totalRegulations int64
	config.DB.Model(&models.Regulation{}).Count(&totalRegulations)

	return c.JSON(fiber.Map{
		"data": categories,
		"meta": fiber.Map{
			"total_categories":  len(categories),
			"total_regulations": totalRegulations,
		},
	})
}

// GetRegulationPageContent returns the regulation page content
func GetRegulationPageContent(c *fiber.Ctx) error {
	var content models.RegulationPageContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.RegulationPageContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}

// getRegulationCategoriesFromTable returns categories from RegulationCategory table with regulation counts
func getRegulationCategoriesFromTable() []fiber.Map {
	var catRows []models.RegulationCategory
	if err := config.DB.Find(&catRows).Error; err != nil {
		return []fiber.Map{}
	}

	categories := make([]fiber.Map, 0, len(catRows))
	for _, cat := range catRows {
		var count int64
		config.DB.Model(&models.Regulation{}).Where("category_id = ?", cat.ID).Count(&count)

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
