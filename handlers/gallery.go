package handlers

import (
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// GetGallery returns gallery images with filtering options
func GetGallery(c *fiber.Ctx) error {
	var images []models.GalleryImage
	query := config.DB.Model(&models.GalleryImage{}).Preload("Category")

	// Apply category filter
	if category := c.Query("category"); category != "" {
		query = query.Where("category_id = ?", category)
	}

	// Apply search filter
	if search := c.Query("search"); search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Apply limit and offset for pagination
	limit := 20 // default
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

	if err := query.Order("sort_order ASC, created_at DESC").Find(&images).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch gallery images",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Get total count
	var total int64
	countQuery := config.DB.Model(&models.GalleryImage{})
	if category := c.Query("category"); category != "" {
		countQuery = countQuery.Where("category_id = ?", category)
	}
	if search := c.Query("search"); search != "" {
		countQuery = countQuery.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	countQuery.Count(&total)

	// Get available categories
	var categories []string
	config.DB.Model(&models.GalleryImage{}).Distinct("category_id").Pluck("category_id", &categories)

	// Calculate pagination
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}
	currentPage := (offset / limit) + 1

	return c.JSON(fiber.Map{
		"data": images,
		"meta": fiber.Map{
			"total":      total,
			"categories": categories,
			"pagination": fiber.Map{
				"current_page": currentPage,
				"per_page":     limit,
				"total_pages":  totalPages,
			},
		},
	})
}

// GetGalleryCategories returns all gallery categories with image counts
func GetGalleryCategories(c *fiber.Ctx) error {
	var categories []models.GalleryCategory

	if err := config.DB.Order("sort_order ASC").Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch gallery categories",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Add count for each category
	result := make([]fiber.Map, len(categories))
	var totalImages int64

	for i, category := range categories {
		var count int64
		config.DB.Model(&models.GalleryImage{}).Where("category_id = ?", category.ID).Count(&count)

		result[i] = fiber.Map{
			"id":             category.ID,
			"name":           category.Name,
			"name_id":        category.NameID,
			"description":    category.Description,
			"description_id": category.DescriptionID,
			"count":          count,
			"color":          category.Color,
			"icon":           category.Icon,
		}

		totalImages += count
	}

	return c.JSON(fiber.Map{
		"data": result,
		"meta": fiber.Map{
			"total_categories": len(categories),
			"total_images":     totalImages,
		},
	})
}

// GetGalleryImageByID returns detailed information about a specific gallery image
func GetGalleryImageByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var image models.GalleryImage
	if err := config.DB.Preload("Category").Where("id = ?", id).First(&image).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Gallery image not found",
			"code":    "NOT_FOUND",
		})
	}

	// Get related images (same category, different images)
	var relatedImages []models.GalleryImage
	config.DB.Where("category_id = ? AND id != ?", image.CategoryID, id).
		Limit(3).
		Find(&relatedImages)

	relatedIDs := make([]string, len(relatedImages))
	for i, img := range relatedImages {
		relatedIDs[i] = img.ID
	}

	response := fiber.Map{
		"id":           image.ID,
		"title":        image.Title,
		"description":  image.Description,
		"image_url":    image.ImageURL,
		"high_res_url": image.HighResURL,
		"metadata": fiber.Map{
			"camera":   "", // You might want to extract this from metadata JSON
			"settings": "", // You might want to extract this from metadata JSON
			"location": image.Location,
		},
		"related_images": relatedIDs,
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}
