package handlers

import (
	"fmt"
	"strconv"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
)

// =============================================================================
// CAROUSEL MANAGEMENT
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

	if err := config.DB.Delete(&models.Carousel{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete carousel slide",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Carousel slide deleted successfully",
	})
}

// =============================================================================
// WHY VISIT MANAGEMENT
// =============================================================================

// CreateWhyVisit creates a new why visit item
func CreateWhyVisit(c *fiber.Ctx) error {
	var whyVisit models.WhyVisit
	if err := c.BodyParser(&whyVisit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create why visit item",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(whyVisit)
}

// UpdateWhyVisit updates an existing why visit item
func UpdateWhyVisit(c *fiber.Ctx) error {
	id := c.Params("id")

	var whyVisit models.WhyVisit
	if err := config.DB.Where("id = ?", id).First(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Why visit item not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&whyVisit); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&whyVisit).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update why visit item",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(whyVisit)
}

// DeleteWhyVisit deletes a why visit item
func DeleteWhyVisit(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Delete(&models.WhyVisit{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete why visit item",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Why visit item deleted successfully",
	})
}

// =============================================================================
// GENERAL WHY VISIT CONTENT MANAGEMENT
// =============================================================================

// UpdateGeneralWhyVisitContent updates the general why visit content (singleton)
func UpdateGeneralWhyVisitContent(c *fiber.Ctx) error {
	var content models.GeneralWhyVisitContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.GeneralWhyVisitContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create general why visit content",
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
			"message": "Failed to update general why visit content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// SELLING POINTS MANAGEMENT
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
// ATTRACTION MANAGEMENT
// =============================================================================

// CreateAttraction creates a new attraction
func CreateAttraction(c *fiber.Ctx) error {
	var attraction models.Attraction
	if err := c.BodyParser(&attraction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Create(&attraction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create attraction",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(attraction)
}

// UpdateAttraction updates an existing attraction
func UpdateAttraction(c *fiber.Ctx) error {
	id := c.Params("id")

	var attraction models.Attraction
	if err := config.DB.Where("id = ?", id).First(&attraction).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Attraction not found",
			"code":    "NOT_FOUND",
		})
	}

	if err := c.BodyParser(&attraction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	if err := config.DB.Save(&attraction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update attraction",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(attraction)
}

// DeleteAttraction deletes an attraction
func DeleteAttraction(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := config.DB.Where("id = ?", id).Delete(&models.Attraction{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to delete attraction",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Attraction deleted successfully",
	})
}

// =============================================================================
// PRICING MANAGEMENT
// =============================================================================

// UpdatePricing updates entrance fee pricing
func UpdatePricing(c *fiber.Ctx) error {
	var pricing models.Pricing
	if err := c.BodyParser(&pricing); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing pricing by type, create if not exists
	var existingPricing models.Pricing
	if err := config.DB.Where("type = ?", pricing.Type).First(&existingPricing).Error; err != nil {
		// Create new pricing
		if err := config.DB.Create(&pricing).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create pricing",
				"code":    "INTERNAL_ERROR",
			})
		}
		return c.Status(fiber.StatusCreated).JSON(pricing)
	}

	// Update existing pricing
	pricing.ID = existingPricing.ID
	if err := config.DB.Save(&pricing).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update pricing",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(pricing)
}

// =============================================================================
// GENERAL WHY VISIT CONTENT MANAGEMENT
// =============================================================================

// UpdateGeneralPricingContent updates the general pricing content (singleton)
func UpdateGeneralPricingContent(c *fiber.Ctx) error {
	var content models.GeneralPricingContent
	if err := c.BodyParser(&content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing content, create if not exists
	var existingContent models.GeneralPricingContent
	if err := config.DB.First(&existingContent).Error; err != nil {
		// Create new content
		if err := config.DB.Create(&content).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create general pricing content",
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
			"message": "Failed to update general pricing content",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(content)
}

// =============================================================================
// PROFILE MANAGEMENT
// =============================================================================

// UpdateProfile updates village profile information
func UpdateProfile(c *fiber.Ctx) error {
	var profile models.Profile
	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Try to find existing profile, create if not exists
	var existingProfile models.Profile
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
// CONTENT UPLOAD
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

// =============================================================================
// CONTACT & BOOKING MANAGEMENT
// =============================================================================

// GetContacts returns all contact form submissions
func GetContacts(c *fiber.Ctx) error {
	var contacts []models.ContactSubmission
	query := config.DB.Model(&models.ContactSubmission{})

	// Apply status filter
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Apply pagination
	limit := 20
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

	if err := query.Limit(limit).Offset(offset).Order("created_at DESC").Find(&contacts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch contact submissions",
			"code":    "INTERNAL_ERROR",
		})
	}

	var total int64
	config.DB.Model(&models.ContactSubmission{}).Count(&total)

	return c.JSON(fiber.Map{
		"data": contacts,
		"meta": fiber.Map{
			"total": total,
			"pagination": fiber.Map{
				"limit":  limit,
				"offset": offset,
			},
		},
	})
}

// GetContactByID returns a specific contact submission
func GetContactByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var contact models.ContactSubmission
	if err := config.DB.Where("id = ?", id).First(&contact).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Contact submission not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": contact,
	})
}

// UpdateContactStatus updates contact submission status
func UpdateContactStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var contact models.ContactSubmission
	if err := config.DB.Where("id = ?", id).First(&contact).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Contact submission not found",
			"code":    "NOT_FOUND",
		})
	}

	var updateData struct {
		Status     string `json:"status"`
		AdminNotes string `json:"admin_notes"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	contact.Status = updateData.Status
	contact.AdminNotes = updateData.AdminNotes

	if err := config.DB.Save(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update contact submission",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(contact)
}

// GetBookings returns all facility bookings
func GetBookings(c *fiber.Ctx) error {
	var bookings []models.Booking
	query := config.DB.Model(&models.Booking{}).Preload("Facility")

	// Apply status filter
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Apply facility filter
	if facilityID := c.Query("facility_id"); facilityID != "" {
		query = query.Where("facility_id = ?", facilityID)
	}

	if err := query.Order("created_at DESC").Find(&bookings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch bookings",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"data": bookings,
	})
}

// UpdateBookingStatus updates booking status
func UpdateBookingStatus(c *fiber.Ctx) error {
	id := c.Params("id")

	var booking models.Booking
	if err := config.DB.Where("id = ?", id).First(&booking).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Booking not found",
			"code":    "NOT_FOUND",
		})
	}

	var updateData struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	booking.Status = updateData.Status

	if err := config.DB.Save(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to update booking status",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(booking)
}
