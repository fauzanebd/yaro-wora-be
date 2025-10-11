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
// GALLERY MANAGEMENT - ADMIN
// =============================================================================

// UpdateGalleryPageContent updates the gallery page content (singleton)
func UpdateGalleryPageContent(c *fiber.Ctx) error {
	var content models.GalleryPageContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.GalleryPageContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create gallery page content",
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
			"message": "Failed to update gallery page content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// CreateGalleryCategory creates a new gallery category
func CreateGalleryCategory(c *fiber.Ctx) error {
	var category models.GalleryCategory
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
			"message": "Failed to create gallery category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// UpdateGalleryCategory updates an existing gallery category
func UpdateGalleryCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.GalleryCategory
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Gallery category not found",
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
			"message": "Failed to update gallery category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(category)
}

// DeleteGalleryCategory deletes a gallery category
func DeleteGalleryCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.GalleryCategory{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete gallery category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Gallery category deleted successfully",
	})
}

// CreateGalleryImage creates a new gallery image
func CreateGalleryImage(c *fiber.Ctx) error {
	var image models.GalleryImage
	if err := c.BodyParser(&image); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&image).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create gallery image",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(image)
}

// UpdateGalleryImage updates an existing gallery image
func UpdateGalleryImage(c *fiber.Ctx) error {
	id := c.Params("id")

	var image models.GalleryImage
	if err := config.DB.Where("id = ?", id).First(&image).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Gallery image not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&image); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&image).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update gallery image",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(image)
}

// DeleteGalleryImage deletes a gallery image
func DeleteGalleryImage(c *fiber.Ctx) error {
	id := c.Params("id")

	// First, get the image to retrieve URLs for R2 deletion
	var image models.GalleryImage
	if err := config.DB.Where("id = ?", id).First(&image).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Gallery image not found",
			"code":    "NOT_FOUND",
		})
	}

	// Delete from database
	if err := config.DB.Delete(&models.GalleryImage{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete gallery image",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Delete images from R2 if they are from our R2 bucket
	if utils.Storage != nil {
		// Delete main image
		if err := utils.Storage.DeleteImageWithThumbnailIfR2(image.ImageURL); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete image from R2: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Gallery image deleted successfully",
	})
}

// =============================================================================
// GALLERY MANAGEMENT - PUBLIC
// =============================================================================

// GetGalleryPageContent returns the gallery page content (public version)
func GetGalleryPageContent(c *fiber.Ctx) error {
	var content models.GalleryPageContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.GalleryPageContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}

// GetGallery returns gallery images with filtering and pagination
func GetGallery(c *fiber.Ctx) error {
	var images []models.GalleryImage
	query := config.DB.Model(&models.GalleryImage{}).
		Select("id, title, title_id, short_description, short_description_id, image_url, thumbnail_url, category_id, date_uploaded").
		Preload("GalleryCategory")

	// Apply category filter (supports comma-separated list)
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

	if err := query.Order("date_uploaded DESC, created_at DESC").Find(&images).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch gallery data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Convert to summary format
	imageSummaries := make([]models.GalleryImageSummary, len(images))
	for i, img := range images {
		imageSummaries[i] = models.GalleryImageSummary{
			ID:                 img.ID,
			Title:              img.Title,
			TitleID:            img.TitleID,
			ShortDescription:   img.ShortDescription,
			ShortDescriptionID: img.ShortDescriptionID,
			ImageURL:           img.ImageURL,
			ThumbnailURL:       img.ThumbnailURL,
			CategoryID:         img.CategoryID,
			GalleryCategory:    img.GalleryCategory,
			DateUploaded:       img.DateUploaded,
		}
	}

	// Get total count (apply same filters)
	var total int64
	countQuery := config.DB.Model(&models.GalleryImage{})
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

	// Get categories with counts from GalleryCategory table
	categories := getGalleryCategoriesFromTable()

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
		"data": imageSummaries,
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

// GetGalleryCategories returns all gallery categories with counts
func GetGalleryCategories(c *fiber.Ctx) error {
	categories := getGalleryCategoriesFromTable()

	var totalImages int64
	config.DB.Model(&models.GalleryImage{}).Count(&totalImages)

	return c.JSON(fiber.Map{
		"data": categories,
		"meta": fiber.Map{
			"total_categories": len(categories),
			"total_images":     totalImages,
		},
	})
}

// GetGalleryImageByID returns a specific gallery image by ID
func GetGalleryImageByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var image models.GalleryImage
	if err := config.DB.Preload("GalleryCategory").Where("id = ?", id).First(&image).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Gallery image not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": image,
	})
}

// getGalleryCategoriesFromTable returns categories from GalleryCategory table with image counts
func getGalleryCategoriesFromTable() []fiber.Map {
	var catRows []models.GalleryCategory
	if err := config.DB.Find(&catRows).Error; err != nil {
		return []fiber.Map{}
	}

	categories := make([]fiber.Map, 0, len(catRows))
	for _, cat := range catRows {
		var count int64
		config.DB.Model(&models.GalleryImage{}).Where("category_id = ?", cat.ID).Count(&count)

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
