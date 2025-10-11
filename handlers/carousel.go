package handlers

import (
	"fmt"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// CAROUSEL MANAGEMENT - ADMIN
// =============================================================================

// CreateCarousel creates a new carousel slide
func CreateCarousel(c *fiber.Ctx) error {
	var carousel models.Carousel
	if err := c.BodyParser(&carousel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&carousel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create carousel slide",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(carousel)
}

// UpdateCarousel updates an existing carousel slide
func UpdateCarousel(c *fiber.Ctx) error {
	id := c.Params("id")

	var carousel models.Carousel
	if err := config.DB.Where("id = ?", id).First(&carousel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Carousel slide not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&carousel); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&carousel).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update carousel slide",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(carousel)
}

// DeleteCarousel deletes a carousel slide
func DeleteCarousel(c *fiber.Ctx) error {
	id := c.Params("id")

	// First, get the carousel to retrieve URLs for R2 deletion
	var carousel models.Carousel
	if err := config.DB.Where("id = ?", id).First(&carousel).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Carousel slide not found",
			"code":    "NOT_FOUND",
		})
	}

	// Delete from database
	if err := config.DB.Delete(&models.Carousel{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete carousel slide",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Delete images from R2 if they are from our R2 bucket
	if utils.Storage != nil {
		// Delete main image
		if err := utils.Storage.DeleteImageWithThumbnailIfR2(carousel.ImageURL); err != nil {
			// Log error but don't fail the request since DB deletion succeeded
			fmt.Printf("Warning: Failed to delete image from R2: %v\n", err)
		}
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Carousel slide deleted successfully",
	})
}

// =============================================================================
// CAROUSEL MANAGEMENT - PUBLIC
// =============================================================================

// GetCarousel returns all active carousel slides
func GetCarousel(c *fiber.Ctx) error {
	var carousels []models.Carousel

	if err := config.DB.Where("is_active = ?", true).Order("carousel_order ASC, created_at ASC").Find(&carousels).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch carousel data",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"data": carousels,
		"meta": fiber.Map{
			"total":              len(carousels),
			"auto_play_interval": 6000,
		},
	})
}
