package handlers

import (
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// GetNews returns news articles with filtering options
func GetNews(c *fiber.Ctx) error {
	var articles []models.NewsArticle
	query := config.DB.Model(&models.NewsArticle{}).Preload("Category")

	// Apply category filter
	if category := c.Query("category"); category != "" {
		query = query.Where("category_key = ?", category)
	}

	// Apply featured filter
	if featured := c.Query("featured"); featured == "true" {
		query = query.Where("is_headline = ?", true)
	}

	// Apply search filter
	if search := c.Query("search"); search != "" {
		query = query.Where("title ILIKE ? OR excerpt ILIKE ? OR content ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Apply language filter
	lang := c.Query("lang", "en")
	query = query.Where("language = ?", lang)

	// Apply pagination
	limit := 12 // default
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

	if err := query.Order("is_headline DESC, date_published DESC").Find(&articles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch news articles",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Get total count
	var total int64
	countQuery := config.DB.Model(&models.NewsArticle{}).Where("language = ?", lang)
	if category := c.Query("category"); category != "" {
		countQuery = countQuery.Where("category_key = ?", category)
	}
	if featured := c.Query("featured"); featured == "true" {
		countQuery = countQuery.Where("is_headline = ?", true)
	}
	if search := c.Query("search"); search != "" {
		countQuery = countQuery.Where("title ILIKE ? OR excerpt ILIKE ? OR content ILIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	countQuery.Count(&total)

	// Get featured count
	var featuredCount int64
	config.DB.Model(&models.NewsArticle{}).Where("is_headline = ? AND language = ?", true, lang).Count(&featuredCount)

	// Get categories with counts
	categories := getNewsCategoriesWithCounts(lang)

	// Calculate pagination
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}
	currentPage := (offset / limit) + 1

	return c.JSON(fiber.Map{
		"data": articles,
		"meta": fiber.Map{
			"total":          total,
			"featured_count": featuredCount,
			"categories":     categories,
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

// GetNewsCategories returns all news categories with article counts
func GetNewsCategories(c *fiber.Ctx) error {
	var categories []models.NewsCategory

	if err := config.DB.Order("sort_order ASC").Find(&categories).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch news categories",
			"code":    "INTERNAL_ERROR",
		})
	}

	lang := c.Query("lang", "en")

	// Add count for each category
	result := make([]fiber.Map, len(categories))
	var totalArticles int64

	for i, category := range categories {
		var count int64
		config.DB.Model(&models.NewsArticle{}).Where("category_key = ? AND language = ?", category.Key, lang).Count(&count)

		result[i] = fiber.Map{
			"key":            category.Key,
			"name":           category.Name,
			"name_id":        category.NameID,
			"description":    category.Description,
			"description_id": category.DescriptionID,
			"count":          count,
			"color":          category.Color,
			"icon":           category.Icon,
		}

		totalArticles += count
	}

	return c.JSON(fiber.Map{
		"data": result,
		"meta": fiber.Map{
			"total_categories": len(categories),
			"total_articles":   totalArticles,
		},
	})
}

// GetNewsByID returns full article content with related articles
func GetNewsByID(c *fiber.Ctx) error {
	id := c.Params("id")
	lang := c.Query("lang", "en")

	var article models.NewsArticle
	if err := config.DB.Preload("Category").Where("id = ? AND language = ?", id, lang).First(&article).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "News article not found",
			"code":    "NOT_FOUND",
		})
	}

	// Increment view count
	config.DB.Model(&article).Update("view_count", article.ViewCount+1)

	// Get related articles (same category, different articles)
	var relatedArticles []models.NewsArticle
	config.DB.Where("category_key = ? AND id != ? AND language = ?", article.CategoryKey, id, lang).
		Select("id, title, excerpt, featured_image, date_published, category_key").
		Limit(3).
		Order("date_published DESC").
		Find(&relatedArticles)

	related := make([]fiber.Map, len(relatedArticles))
	for i, rel := range relatedArticles {
		related[i] = fiber.Map{
			"id":             rel.ID,
			"title":          rel.Title,
			"excerpt":        rel.Excerpt,
			"featured_image": rel.FeaturedImage,
			"date_published": rel.DatePublished,
			"category":       rel.CategoryKey,
		}
	}

	// Build author social links
	authorSocial := make(map[string]interface{})
	if len(article.AuthorSocial) > 0 {
		// AuthorSocial is datatypes.JSON, we can use it directly or unmarshal it
		authorSocial = make(map[string]interface{})
	}

	// Build SEO data
	seoKeywords := make([]string, 0)
	if len(article.SEOKeywords) > 0 {
		// SEOKeywords is datatypes.JSON, we can use it directly or unmarshal it
		seoKeywords = make([]string, 0)
	}

	response := fiber.Map{
		"id":         article.ID,
		"title":      article.Title,
		"title_id":   article.TitleID,
		"excerpt":    article.Excerpt,
		"excerpt_id": article.ExcerptID,
		"content":    article.Content,
		"content_id": article.ContentID,
		"author": fiber.Map{
			"name":         article.AuthorName,
			"avatar":       article.AuthorAvatar,
			"bio":          article.AuthorBio,
			"email":        article.AuthorEmail,
			"social_links": authorSocial,
		},
		"date_published":   article.DatePublished,
		"date_updated":     article.UpdatedAt,
		"category":         article.CategoryKey,
		"featured_image":   article.FeaturedImage,
		"image_gallery":    article.ImageGallery,
		"tags":             article.Tags,
		"read_time":        article.ReadTime,
		"is_headline":      article.IsHeadline,
		"view_count":       article.ViewCount + 1, // Return incremented count
		"related_articles": related,
		"seo": fiber.Map{
			"meta_title":       article.SEOMetaTitle,
			"meta_description": article.SEOMetaDesc,
			"keywords":         seoKeywords,
			"canonical_url":    article.CanonicalURL,
		},
		"language": article.Language,
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

// GetFeaturedNews returns the current featured/headline article
func GetFeaturedNews(c *fiber.Ctx) error {
	lang := c.Query("lang", "en")

	var article models.NewsArticle
	if err := config.DB.Where("is_headline = ? AND language = ?", true, lang).
		Order("date_published DESC").
		First(&article).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Featured article not found",
			"code":    "NOT_FOUND",
		})
	}

	response := fiber.Map{
		"id":             article.ID,
		"title":          article.Title,
		"excerpt":        article.Excerpt,
		"featured_image": article.FeaturedImage,
		"author": fiber.Map{
			"name":   article.AuthorName,
			"avatar": article.AuthorAvatar,
		},
		"date_published": article.DatePublished,
		"category":       article.CategoryKey,
		"read_time":      article.ReadTime,
		"view_count":     article.ViewCount,
	}

	return c.JSON(fiber.Map{
		"data": response,
	})
}

// Helper function to get news categories with counts
func getNewsCategoriesWithCounts(lang string) []fiber.Map {
	type CategoryCount struct {
		Key   string
		Name  string
		Count int64
	}

	var results []CategoryCount
	config.DB.Table("news_categories").
		Select("news_categories.key, news_categories.name, COUNT(news_articles.id) as count").
		Joins("LEFT JOIN news_articles ON news_categories.key = news_articles.category_key AND news_articles.language = ?", lang).
		Group("news_categories.key, news_categories.name").
		Order("news_categories.sort_order ASC").
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
