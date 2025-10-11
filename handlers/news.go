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
// NEWS MANAGEMENT - ADMIN
// =============================================================================

// CreateNews creates a new news article
func CreateNews(c *fiber.Ctx) error {
	var news models.NewsArticle
	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create news article",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(news)
}

// UpdateNews updates an existing news article
func UpdateNews(c *fiber.Ctx) error {
	id := c.Params("id")

	var news models.NewsArticle
	if err := config.DB.Where("id = ?", id).First(&news).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News article not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update news article",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(news)
}

// DeleteNews deletes a news article
func DeleteNews(c *fiber.Ctx) error {
	id := c.Params("id")

	// First, get the news article to retrieve URLs for R2 deletion
	var news models.NewsArticle
	if err := config.DB.Where("id = ?", id).First(&news).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News article not found",
			"code":    "NOT_FOUND",
		})
	}

	// Delete from database
	if err := config.DB.Delete(&models.NewsArticle{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete news article",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Delete images from R2 if they are from our R2 bucket
	if utils.Storage != nil {
		// Delete main image
		if err := utils.Storage.DeleteImageIfR2(news.ImageURL); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete image from R2: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "News article deleted successfully",
	})
}

// UpdateNewsPageContent updates the news page content (singleton)
func UpdateNewsPageContent(c *fiber.Ctx) error {
	var content models.NewsPageContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.NewsPageContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create news page content",
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
			"message": "Failed to update news page content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// NEWS CATEGORY MANAGEMENT - ADMIN
// =============================================================================

// CreateNewsCategory creates a new news category
func CreateNewsCategory(c *fiber.Ctx) error {
	var category models.NewsCategory
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
			"message": "Failed to create news category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// UpdateNewsCategory updates an existing news category
func UpdateNewsCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.NewsCategory
	if err := config.DB.Where("id = ?", id).First(&category).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News category not found",
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
			"message": "Failed to update news category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(category)
}

// DeleteNewsCategory deletes a news category
func DeleteNewsCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.NewsCategory{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete news category",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "News category deleted successfully",
	})
}

// =============================================================================
// NEWS AUTHOR MANAGEMENT - ADMIN
// =============================================================================

// CreateNewsAuthor creates a new news author
func CreateNewsAuthor(c *fiber.Ctx) error {
	var author models.NewsAuthor
	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create news author",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(author)
}

// UpdateNewsAuthor updates an existing news author
func UpdateNewsAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	var author models.NewsAuthor
	if err := config.DB.Where("id = ?", id).First(&author).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News author not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&author); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&author).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update news author",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(author)
}

// DeleteNewsAuthor deletes a news author
func DeleteNewsAuthor(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.NewsAuthor{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete news author",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "News author deleted successfully",
	})
}

// =============================================================================
// NEWS MANAGEMENT - PUBLIC
// =============================================================================

// GetNews returns all news articles with filtering options
func GetNews(c *fiber.Ctx) error {
	var news []models.NewsArticle
	query := config.DB.Model(&models.NewsArticle{}).
		Select("id, title, title_id, excerpt, excerpt_id, image_url, date_published, author_id, category_id, tags, read_time, is_headline").
		Preload("NewsAuthor").
		Preload("NewsCategory")

	// Apply headline filter
	isHeadline := c.Query("headline")
	if isHeadline != "" {
		if isHeadline == "true" {
			query = query.Where("is_headline = ?", true)
		} else {
			query = query.Where("is_headline = ?", false)
		}
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

	// Apply author filter
	if author := c.Query("author"); author != "" {
		query = query.Where("author_id = ?", author)
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

	if err := query.Order("date_published DESC, created_at DESC").Find(&news).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch news data",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Convert to summary format
	newsSummaries := make([]models.NewsArticleSummary, len(news))
	for i, article := range news {
		newsSummaries[i] = models.NewsArticleSummary{
			ID:            article.ID,
			Title:         article.Title,
			TitleID:       article.TitleID,
			Excerpt:       article.Excerpt,
			ExcerptID:     article.ExcerptID,
			AuthorID:      article.AuthorID,
			NewsAuthor:    article.NewsAuthor,
			DatePublished: article.DatePublished,
			CategoryID:    article.CategoryID,
			NewsCategory:  article.NewsCategory,
			ImageURL:      article.ImageURL,
			Tags:          article.Tags,
			ReadTime:      article.ReadTime,
			IsHeadline:    article.IsHeadline,
		}
	}

	// Get total count (apply same filters)
	var total int64
	countQuery := config.DB.Model(&models.NewsArticle{})
	if isHeadline := c.Query("headline"); isHeadline == "true" {
		countQuery = countQuery.Where("is_headline = ?", true)
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
	if author := c.Query("author"); author != "" {
		countQuery = countQuery.Where("author_id = ?", author)
	}
	countQuery.Count(&total)

	// Get categories with counts from NewsCategory table
	categories := getNewsCategoriesFromTable()

	// Get authors with counts from NewsAuthor table
	authors := getNewsAuthorsFromTable()

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
		"data": newsSummaries,
		"meta": fiber.Map{
			"total":      total,
			"categories": categories,
			"authors":    authors,
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

// GetNewsByID returns a specific news article by ID
func GetNewsByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var news models.NewsArticle
	if err := config.DB.Preload("NewsAuthor").Preload("NewsCategory").Where("id = ?", id).First(&news).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News article not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": news,
	})
}

// GetNewsCategories returns all news categories with counts
func GetNewsCategories(c *fiber.Ctx) error {
	categories := getNewsCategoriesFromTable()

	var totalNews int64
	config.DB.Model(&models.NewsArticle{}).Count(&totalNews)

	return c.JSON(fiber.Map{
		"data": categories,
		"meta": fiber.Map{
			"total_categories": len(categories),
			"total_news":       totalNews,
		},
	})
}

// GetNewsAuthors returns all news authors with counts
func GetNewsAuthors(c *fiber.Ctx) error {
	authors := getNewsAuthorsFromTable()

	var totalNews int64
	config.DB.Model(&models.NewsArticle{}).Count(&totalNews)

	return c.JSON(fiber.Map{
		"data": authors,
		"meta": fiber.Map{
			"total_authors": len(authors),
			"total_news":    totalNews,
		},
	})
}

// GetNewsAuthorByID returns a specific news author by ID
func GetNewsAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var author models.NewsAuthor
	if err := config.DB.Where("id = ?", id).First(&author).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News author not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": author,
	})
}

// GetNewsPageContent returns the news page content
func GetNewsPageContent(c *fiber.Ctx) error {
	var content models.NewsPageContent
	if err := config.DB.First(&content).Error; err != nil {
		// Return empty content if not found
		return c.JSON(fiber.Map{
			"data": models.NewsPageContent{},
		})
	}
	return c.JSON(fiber.Map{
		"data": content,
	})
}

// getNewsCategoriesFromTable returns categories from NewsCategory table with news counts
func getNewsCategoriesFromTable() []fiber.Map {
	var catRows []models.NewsCategory
	if err := config.DB.Find(&catRows).Error; err != nil {
		return []fiber.Map{}
	}

	categories := make([]fiber.Map, 0, len(catRows))
	for _, cat := range catRows {
		var count int64
		config.DB.Model(&models.NewsArticle{}).Where("category_id = ?", cat.ID).Count(&count)

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

// getNewsAuthorsFromTable returns authors from NewsAuthor table with news counts
func getNewsAuthorsFromTable() []fiber.Map {
	var authorRows []models.NewsAuthor
	if err := config.DB.Find(&authorRows).Error; err != nil {
		return []fiber.Map{}
	}

	authors := make([]fiber.Map, 0, len(authorRows))
	for _, author := range authorRows {
		var count int64
		config.DB.Model(&models.NewsArticle{}).Where("author_id = ?", author.ID).Count(&count)

		authors = append(authors, fiber.Map{
			"id":     author.ID,
			"name":   author.Name,
			"avatar": author.Avatar,
			"count":  count,
		})
	}
	return authors
}
