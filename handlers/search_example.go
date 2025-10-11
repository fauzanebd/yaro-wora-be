package handlers

import (
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// SearchUsers demonstrates case-insensitive user search with citext
func SearchUsers(c *fiber.Ctx) error {
	searchTerm := c.Query("q", "")

	var users []models.User
	query := config.DB.Model(&models.User{})

	if searchTerm != "" {
		// With citext username field, this automatically works case-insensitively
		// So "ADMIN", "admin", "Admin" will all match the same record
		query = query.Where("username = ? OR username ILIKE ?", searchTerm, "%"+searchTerm+"%")
	}

	if err := query.Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to search users",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"total":       len(users),
			"search_term": searchTerm,
		},
	})
}

// // SearchContactsByEmail demonstrates citext email search
// func SearchContactsByEmail(c *fiber.Ctx) error {
// 	email := c.Query("email", "")

// 	if email == "" {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Email parameter is required",
// 			"code":    "BAD_REQUEST",
// 		})
// 	}

// 	var contacts []models.ContactSubmission

// 	// With citext email field, this automatically works case-insensitively
// 	// "TEST@EXAMPLE.COM", "test@example.com", "Test@Example.com" all match
// 	if err := utils.Search.EmailSearch(config.DB, email).Find(&contacts).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Failed to search contacts",
// 			"code":    "INTERNAL_ERROR",
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"data": contacts,
// 		"meta": fiber.Map{
// 			"total": len(contacts),
// 			"email": email,
// 		},
// 	})
// }

// AdvancedContentSearch demonstrates multi-field search
func AdvancedContentSearch(c *fiber.Ctx) error {
	searchTerm := c.Query("q", "")
	contentType := c.Query("type", "all") // news, destinations, facilities
	exactMatch := c.Query("exact", "false") == "true"

	limit := 20
	if l := c.Query("limit"); l != "" {
		if limitInt, err := strconv.Atoi(l); err == nil && limitInt > 0 {
			limit = limitInt
		}
	}

	results := make(map[string]interface{})

	if contentType == "all" || contentType == "news" {
		var news []models.NewsArticle
		searchConfig := utils.SearchConfig{
			Query:      searchTerm,
			Fields:     []string{"title", "excerpt", "content"},
			ExactMatch: exactMatch,
		}

		query := utils.Search.AdvancedSearch(config.DB.Model(&models.NewsArticle{}), searchConfig)
		query.Limit(limit).Find(&news)
		results["news"] = news
	}

	if contentType == "all" || contentType == "destinations" {
		var destinations []models.Destination
		searchConfig := utils.SearchConfig{
			Query:      searchTerm,
			Fields:     []string{"title", "description", "location_address"},
			ExactMatch: exactMatch,
		}

		query := utils.Search.AdvancedSearch(config.DB.Model(&models.Destination{}), searchConfig)
		query.Limit(limit).Find(&destinations)
		results["destinations"] = destinations
	}

	if contentType == "all" || contentType == "facilities" {
		var facilities []models.Facility
		searchConfig := utils.SearchConfig{
			Query:      searchTerm,
			Fields:     []string{"name", "description", "location_address"},
			ExactMatch: exactMatch,
		}

		query := utils.Search.AdvancedSearch(config.DB.Model(&models.Facility{}), searchConfig)
		query.Limit(limit).Find(&facilities)
		results["facilities"] = facilities
	}

	return c.JSON(fiber.Map{
		"data": results,
		"meta": fiber.Map{
			"search_term":  searchTerm,
			"content_type": contentType,
			"exact_match":  exactMatch,
			"limit":        limit,
		},
	})
}

// GetUserByUsernameOrEmail demonstrates citext field querying
func GetUserByUsernameOrEmail(c *fiber.Ctx) error {
	identifier := c.Params("identifier") // Can be username or email

	var user models.User

	// Check if it looks like an email
	if utils.Search.ValidateEmail(identifier) {
		// Search by email if it looks like one
		// Note: We don't have email in User model, but this shows the pattern
		err := config.DB.Where("username = ?", identifier).First(&user).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "User not found",
				"code":    "NOT_FOUND",
			})
		}
	} else {
		// Search by username (citext field - case insensitive)
		err := utils.Search.UsernameSearch(config.DB, identifier).First(&user).Error
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "User not found",
				"code":    "NOT_FOUND",
			})
		}
	}

	return c.JSON(fiber.Map{
		"data": user,
	})
}
