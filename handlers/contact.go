package handlers

import (
	"encoding/json"
	"fmt"
	"time"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// ContactRequest represents the contact form submission
type ContactRequest struct {
	Name          string         `json:"name" validate:"required"`
	Email         string         `json:"email" validate:"required,email"`
	Phone         string         `json:"phone"`
	Subject       string         `json:"subject" validate:"required"`
	Message       string         `json:"message" validate:"required"`
	PreferredDate string         `json:"preferred_date"`
	VisitorType   string         `json:"visitor_type"`  // domestic, locals_sumba, foreigner
	VisitorCount  map[string]int `json:"visitor_count"` // {adults: int, infants: int}
}

// SubmitContact handles contact form submissions
func SubmitContact(c *fiber.Ctx) error {
	var req ContactRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Subject == "" || req.Message == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Name, email, subject, and message are required",
			"code":    "BAD_REQUEST",
		})
	}

	// Generate reference ID
	referenceID := fmt.Sprintf("INQ-%s-%s",
		time.Now().Format("2006"),
		uuid.New().String()[:8])

	// Convert visitor count to JSON
	var visitorCountJSON datatypes.JSON
	if req.VisitorCount != nil {
		// Convert map to JSON bytes, then to datatypes.JSON
		if jsonBytes, err := json.Marshal(req.VisitorCount); err == nil {
			visitorCountJSON = datatypes.JSON(jsonBytes)
		}
	}

	// Create contact submission
	submission := models.ContactSubmission{
		Name:          req.Name,
		Email:         req.Email,
		Phone:         req.Phone,
		Subject:       req.Subject,
		Message:       req.Message,
		PreferredDate: req.PreferredDate,
		VisitorType:   req.VisitorType,
		VisitorCount:  visitorCountJSON,
		Status:        "pending",
		ReferenceID:   referenceID,
		ResponseSent:  false,
	}

	if err := config.DB.Create(&submission).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to submit contact form",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":                 true,
		"message":                 "Your inquiry has been submitted successfully. We will contact you within 24 hours.",
		"reference_id":            referenceID,
		"estimated_response_time": "24 hours",
	})
}

// GetContactInfo returns contact information and location details
func GetContactInfo(c *fiber.Ctx) error {
	var contactInfo models.ContactInfo

	if err := config.DB.First(&contactInfo).Error; err != nil {
		// Return default contact info if not found in database
		return c.JSON(fiber.Map{
			"data": fiber.Map{
				"address": fiber.Map{
					"street":      "Yaro Wora Village",
					"city":        "East Sumba",
					"province":    "East Nusa Tenggara",
					"country":     "Indonesia",
					"postal_code": "87173",
				},
				"coordinates": fiber.Map{
					"latitude":  -9.6234,
					"longitude": 119.3456,
				},
				"contact": fiber.Map{
					"phones":   []string{"+62 098 940 974", "+62 903 009 909"},
					"email":    []string{"info@yarowora.com", "visit@yarowora.com"},
					"whatsapp": "+62 903 009 909",
				},
				"social_media": fiber.Map{
					"instagram": "@yarowora_official",
					"facebook":  "Yaro Wora Tourism",
					"youtube":   "Yaro Wora Channel",
				},
				"operating_hours": fiber.Map{
					"monday":    "08:00-17:00",
					"tuesday":   "08:00-17:00",
					"wednesday": "08:00-17:00",
					"thursday":  "08:00-17:00",
					"friday":    "08:00-17:00",
					"saturday":  "08:00-16:00",
					"sunday":    "closed",
				},
			},
		})
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"address": fiber.Map{
				"street":      contactInfo.Street,
				"city":        contactInfo.City,
				"province":    contactInfo.Province,
				"country":     contactInfo.Country,
				"postal_code": contactInfo.PostalCode,
			},
			"coordinates": fiber.Map{
				"latitude":  contactInfo.Latitude,
				"longitude": contactInfo.Longitude,
			},
			"contact": fiber.Map{
				"phones":   contactInfo.Phones,
				"email":    contactInfo.Emails,
				"whatsapp": contactInfo.WhatsApp,
			},
			"social_media":    contactInfo.SocialMedia,
			"operating_hours": contactInfo.OperatingHours,
		},
	})
}
