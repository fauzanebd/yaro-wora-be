package handlers

import (
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// SELLING POINTS MANAGEMENT - ADMIN
// =============================================================================

// CreateSellingPoint creates a new selling point
func CreateSellingPoint(c *fiber.Ctx) error {
	var sellingPoint models.SellingPoint
	if err := c.BodyParser(&sellingPoint); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&sellingPoint).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create selling point",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(sellingPoint)
}

// UpdateSellingPoint updates an existing selling point
func UpdateSellingPoint(c *fiber.Ctx) error {
	id := c.Params("id")

	var sellingPoint models.SellingPoint
	if err := config.DB.Where("id = ?", id).First(&sellingPoint).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Selling point not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&sellingPoint); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&sellingPoint).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update selling point",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(sellingPoint)
}

// DeleteSellingPoint deletes a selling point
func DeleteSellingPoint(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.SellingPoint{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete selling point",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Selling point deleted successfully",
	})
}

// =============================================================================
// SELLING POINTS MANAGEMENT - PUBLIC
// =============================================================================

// GetSellingPoints returns all active selling points
func GetSellingPoints(c *fiber.Ctx) error {
	var sellingPoints []models.SellingPoint

	if err := config.DB.Where("is_active = ?", true).Order("selling_point_order ASC, created_at ASC").Find(&sellingPoints).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch selling points data",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"data": sellingPoints,
		"meta": fiber.Map{
			"total": len(sellingPoints),
		},
	})
}
