package handlers

import (
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// PROFILE MANAGEMENT - ADMIN
// =============================================================================

// UpdateProfilePageContent updates village profile information
func UpdateProfilePageContent(c *fiber.Ctx) error {
	var profile models.ProfilePageContent
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing profile, create if not exists
	var existingProfile models.ProfilePageContent
	if err := config.DB.First(&existingProfile).Error; err != nil {
		// Create new profile
		if err := config.DB.Create(&profile).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create profile",
				"code":    "INTERNAL_ERROR",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(profile)
	}

	// Update existing profile
	profile.ID = existingProfile.ID
	if err := config.DB.Save(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update profile",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(profile)
}

// =============================================================================
// PROFILE MANAGEMENT - PUBLIC
// =============================================================================

// GetProfilePageContent returns village profile page content
func GetProfilePageContent(c *fiber.Ctx) error {
	var profile models.ProfilePageContent

	if err := config.DB.First(&profile).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Profile page content not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": profile,
	})
}
