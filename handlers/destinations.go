package handlers

import (
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// GetDestinations returns all destinations with filtering options
func GetDestinations(c *fiber.Ctx) error {
	var destinations []models.Destination
	query := config.DB.Model(&models.Destination{})

	// Apply type filter
	if destType := c.Query("type"); destType != "" {
		query = query.Where("type = ?", destType)
	}

	// Apply category filter
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	// Apply limit and offset for pagination
	if limit := c.Query("limit"); limit != "" {
		if limitInt, err := strconv.Atoi(limit); err == nil && limitInt > 0 {
			query = query.Limit(limitInt)
		}
	}

	if offset := c.Query("offset"); offset != "" {
		if offsetInt, err := strconv.Atoi(offset); err == nil && offsetInt >= 0 {
			query = query.Offset(offsetInt)
		}
	}

	if err := query.Order("sort_order ASC, created_at ASC").Find(&destinations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch destinations data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Get total count
	var total int64
	config.DB.Model(&models.Destination{}).Count(&total)

	// Count main and other destinations
	var mainCount, otherCount int64
	config.DB.Model(&models.Destination{}).Where("type = ?", "main").Count(&mainCount)
	config.DB.Model(&models.Destination{}).Where("type = ?", "other").Count(&otherCount)

	// Get categories with counts
	categories := getDestinationCategories()

	return c.JSON(fiber.Map{
		"data": destinations,
		"meta": fiber.Map{
			"total":              total,
			"main_destinations":  mainCount,
			"other_destinations": otherCount,
			"categories":         categories,
			"pagination": fiber.Map{
				"current_page": 1, // You'll need to calculate this based on offset/limit
				"per_page":     len(destinations),
				"total_pages":  1, // Calculate based on total/limit
			},
		},
	})
}

// GetMainDestination returns the main featured destination
func GetMainDestination(c *fiber.Ctx) error {
	var destination models.Destination

	if err := config.DB.Where("type = ?", "main").First(&destination).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Main destination not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": destination,
	})
}

// GetDestinationByID returns a specific destination by ID
func GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var destination models.Destination
	if err := config.DB.Where("id = ?", id).First(&destination).Error; err != nil {
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

// getDestinationCategories returns available categories with counts
func getDestinationCategories() []fiber.Map {
	type CategoryCount struct {
		Category string
		Count    int64
	}

	var results []CategoryCount
	config.DB.Model(&models.Destination{}).
		Select("category, COUNT(*) as count").
		Group("category").
		Scan(&results)

	categories := make([]fiber.Map, 0)
	categoryNames := map[string]string{
		"nature":      "Nature",
		"culture":     "Cultural Heritage",
		"heritage":    "Heritage",
		"agriculture": "Agriculture",
		"adventure":   "Adventure",
	}

	for _, result := range results {
		name, exists := categoryNames[result.Category]
		if !exists {
			name = result.Category
		}

		categories = append(categories, fiber.Map{
			"id":    result.Category,
			"name":  name,
			"count": result.Count,
		})
	}

	return categories
}
