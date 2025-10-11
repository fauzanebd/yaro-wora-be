package handlers

import (
	"fmt"
	"yaro-wora-be/config"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// CONTENT UPLOAD - ADMIN
// =============================================================================

// UploadContent handles file uploads to Cloudflare R2
func UploadContent(c *fiber.Ctx) error {
	// Get uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "No file uploaded",
			"code":    "BAD_REQUEST",
		})
	}

	// Validate file size
	maxSize := int64(config.AppConfig.MaxFileUploadSize)
	if file.Size > maxSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("File size exceeds maximum limit of %d MB", maxSize/(1024*1024)),
			"code":    "FILE_TOO_LARGE",
		})
	}

	// Check storage limit before upload
	if err := utils.Storage.CheckStorageLimit(file.Size); err != nil {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
			"code":    "STORAGE_LIMIT_EXCEEDED",
		})
	}

	// Determine upload folder based on content type or form field
	folder := c.FormValue("folder", "uploads")

	// Upload to R2 with thumbnail generation and dimension detection
	uploadResponse, err := utils.Storage.UploadImageWithThumbnail(file, folder)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to upload file: " + err.Error(),
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(uploadResponse)
}
