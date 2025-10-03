package handlers

import (
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// GetRegulations returns regulations with filtering options
func GetRegulations(c *fiber.Ctx) error {
	var regulations []models.Regulation
	query := config.DB.Model(&models.Regulation{}).Preload("Category")

	// Apply active filter (default to true)
	isActive := c.Query("is_active", "true")
	if isActive == "true" {
		query = query.Where("is_active = ?", true)
	} else if isActive == "false" {
		query = query.Where("is_active = ?", false)
	}

	// Apply category filter
	if category := c.Query("category"); category != "" {
		query = query.Where("category_key = ?", category)
	}

	// Apply search filter
	if search := c.Query("search"); search != "" {
		query = query.Where("question ILIKE ? OR answer ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Apply pagination
	limit := 20 // default
	if l := c.Query("limit"); l != "" {
		if limitInt, err := strconv.Atoi(l); err == nil && limitInt > 0 && limitInt <= 100 {
			limit = limitInt
		}
	}

	page := 1
	if p := c.Query("page"); p != "" {
		if pageInt, err := strconv.Atoi(p); err == nil && pageInt > 0 {
			page = pageInt
		}
	}

	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	if err := query.Order("priority ASC, created_at DESC").Find(&regulations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch regulations",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Get total count for pagination
	var total int64
	countQuery := config.DB.Model(&models.Regulation{})
	if isActive == "true" {
		countQuery = countQuery.Where("is_active = ?", true)
	} else if isActive == "false" {
		countQuery = countQuery.Where("is_active = ?", false)
	}
	if category := c.Query("category"); category != "" {
		countQuery = countQuery.Where("category_key = ?", category)
	}
	if search := c.Query("search"); search != "" {
		countQuery = countQuery.Where("question ILIKE ? OR answer ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	countQuery.Count(&total)

	// Get categories with counts
	categories := getRegulationCategoriesWithCounts()

	// Calculate pagination
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}

	return c.JSON(fiber.Map{
		"data": regulations,
		"meta": fiber.Map{
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
			"categories":  categories,
		},
	})
}

// GetRegulationCategories returns all regulation categories with statistics
func GetRegulationCategories(c *fiber.Ctx) error {
	var categories []models.RegulationCategory

	if err := config.DB.Order("sort_order ASC").Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch regulation categories",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Add count for each category
	result := make([]fiber.Map, len(categories))
	var totalRegulations int64

	for i, category := range categories {
		var count int64
		config.DB.Model(&models.Regulation{}).Where("category_key = ? AND is_active = ?", category.Key, true).Count(&count)

		result[i] = fiber.Map{
			"key":            category.Key,
			"name":           category.Name,
			"name_id":        category.NameID,
			"description":    category.Description,
			"description_id": category.DescriptionID,
			"count":          count,
			"icon":           category.Icon,
			"color":          category.Color,
		}

		totalRegulations += count
	}

	return c.JSON(fiber.Map{
		"data": result,
		"meta": fiber.Map{
			"total_categories":  len(categories),
			"total_regulations": totalRegulations,
		},
	})
}

// GetRegulationByID returns a specific regulation by ID
func GetRegulationByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var regulation models.Regulation
	if err := config.DB.Preload("Category").Where("id = ? AND is_active = ?", id, true).First(&regulation).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Regulation not found",
			"code":    "NOT_FOUND",
		})
	}

	// Get related regulations (same category, different regulations)
	var relatedRegulations []models.Regulation
	config.DB.Where("category_key = ? AND id != ? AND is_active = ?", regulation.CategoryKey, id, true).
		Select("id, question").
		Limit(3).
		Find(&relatedRegulations)

	related := make([]fiber.Map, len(relatedRegulations))
	for i, rel := range relatedRegulations {
		related[i] = fiber.Map{
			"id":       rel.ID,
			"question": rel.Question,
		}
	}

	response := fiber.Map{
		"id":                  regulation.ID,
		"category":            regulation.CategoryKey,
		"question":            regulation.Question,
		"answer":              regulation.Answer,
		"priority":            regulation.Priority,
		"is_active":           regulation.IsActive,
		"tags":                regulation.Tags,
		"related_regulations": related,
		"created_at":          regulation.CreatedAt,
		"updated_at":          regulation.UpdatedAt,
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

// Helper function to get regulation categories with counts
func getRegulationCategoriesWithCounts() []fiber.Map {
	type CategoryCount struct {
		Key   string
		Name  string
		Count int64
	}

	var results []CategoryCount
	config.DB.Table("regulation_categories").
		Select("regulation_categories.key, regulation_categories.name, COUNT(regulations.id) as count").
		Joins("LEFT JOIN regulations ON regulation_categories.key = regulations.category_key AND regulations.is_active = true").
		Group("regulation_categories.key, regulation_categories.name").
		Order("regulation_categories.sort_order ASC").
		Scan(&results)

	categories := make([]fiber.Map, len(results))
	for i, result := range results {
		categories[i] = fiber.Map{
			"key":   result.Key,
			"name":  result.Name,
			"count": result.Count,
		}
	}

	return categories
}
