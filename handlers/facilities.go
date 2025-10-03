package handlers

import (
	"fmt"
	"strconv"
	"time"
	"yaro-wora-be/config"
	"yaro-wora-be/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetFacilities returns facilities with filtering options
func GetFacilities(c *fiber.Ctx) error {
	var facilities []models.Facility
	query := config.DB.Model(&models.Facility{})

	// Apply category filter
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}

	// Apply sorting
	sortBy := c.Query("sort", "sort_order")
	switch sortBy {
	case "price":
		query = query.Order("price ASC")
	case "duration":
		query = query.Order("duration ASC")
	case "popularity":
		query = query.Order("sort_order ASC") // Assuming sort_order represents popularity
	default:
		query = query.Order("sort_order ASC, created_at DESC")
	}

	// Apply pagination
	limit := 20 // default
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

	if err := query.Find(&facilities).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to fetch facilities",
			"code":    "INTERNAL_ERROR",
		})
	}

	// Get total count
	var total int64
	countQuery := config.DB.Model(&models.Facility{})
	if category := c.Query("category"); category != "" {
		countQuery = countQuery.Where("category = ?", category)
	}
	countQuery.Count(&total)

	// Get categories with counts
	categories := getFacilityCategoriesWithCounts()

	// Calculate pagination
	totalPages := int(total) / limit
	if int(total)%limit != 0 {
		totalPages++
	}
	currentPage := (offset / limit) + 1

	return c.JSON(fiber.Map{
		"data": facilities,
		"meta": fiber.Map{
			"total":      total,
			"categories": categories,
			"pagination": fiber.Map{
				"current_page": currentPage,
				"per_page":     limit,
				"total_pages":  totalPages,
			},
		},
	})
}

// GetFacilityByID returns detailed information about a specific facility
func GetFacilityByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var facility models.Facility
	if err := config.DB.Where("id = ?", id).First(&facility).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Facility not found",
			"code":    "NOT_FOUND",
		})
	}

	return c.JSON(fiber.Map{
		"data": facility,
	})
}

// BookFacility creates a new booking for a facility
func BookFacility(c *fiber.Ctx) error {
	facilityID := c.Params("id")

	// Check if facility exists
	var facility models.Facility
	if err := config.DB.Where("id = ?", facilityID).First(&facility).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Facility not found",
			"code":    "NOT_FOUND",
		})
	}

	// Parse request body
	var req BookingRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request body",
			"code":    "BAD_REQUEST",
		})
	}

	// Validate required fields
	if req.GuestName == "" || req.Email == "" || req.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Guest name, email, and phone are required",
			"code":    "BAD_REQUEST",
		})
	}

	// Calculate total price (simple calculation based on participants)
	totalPrice := facility.Price * req.Participants
	if facility.Category == "accommodation" && req.CheckOutDate != "" {
		// Calculate nights for accommodation
		checkIn, _ := time.Parse("2006-01-02", req.CheckInDate)
		checkOut, _ := time.Parse("2006-01-02", req.CheckOutDate)
		nights := int(checkOut.Sub(checkIn).Hours() / 24)
		if nights > 0 {
			totalPrice = facility.Price * nights
		}
	}

	// Generate booking ID
	bookingID := fmt.Sprintf("BOOK-%s-%s",
		time.Now().Format("2006"),
		uuid.New().String()[:8])

	// Create booking
	booking := models.Booking{
		BookingID:             bookingID,
		FacilityID:            facilityID,
		GuestName:             req.GuestName,
		Email:                 req.Email,
		Phone:                 req.Phone,
		CheckInDate:           req.CheckInDate,
		CheckOutDate:          req.CheckOutDate,
		Participants:          req.Participants,
		SpecialRequirements:   req.SpecialRequirements,
		LanguagePreference:    req.LanguagePreference,
		TotalPrice:            totalPrice,
		Currency:              facility.Currency,
		Status:                "pending",
		ConfirmationEmailSent: true, // You might want to implement actual email sending
	}

	if err := config.DB.Create(&booking).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to create booking",
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":                 true,
		"booking_id":              bookingID,
		"message":                 "Booking confirmed successfully",
		"total_price":             totalPrice,
		"currency":                facility.Currency,
		"confirmation_email_sent": true,
	})
}

// BookingRequest represents the booking request body
type BookingRequest struct {
	GuestName           string `json:"guest_name"`
	Email               string `json:"email"`
	Phone               string `json:"phone"`
	CheckInDate         string `json:"check_in_date"`
	CheckOutDate        string `json:"check_out_date"`
	Participants        int    `json:"participants"`
	SpecialRequirements string `json:"special_requirements"`
	LanguagePreference  string `json:"language_preference"`
}

// Helper function to get facility categories with counts
func getFacilityCategoriesWithCounts() []fiber.Map {
	type CategoryCount struct {
		Category string
		Count    int64
	}

	var results []CategoryCount
	config.DB.Model(&models.Facility{}).
		Select("category, COUNT(*) as count").
		Group("category").
		Scan(&results)

	categoryNames := map[string]map[string]string{
		"accommodation": {"name": "Accommodation", "name_id": "Akomodasi"},
		"workshop":      {"name": "Workshop", "name_id": "Workshop"},
		"culinary":      {"name": "Culinary", "name_id": "Kuliner"},
		"entertainment": {"name": "Entertainment", "name_id": "Hiburan"},
		"activity":      {"name": "Activity", "name_id": "Aktivitas"},
		"wellness":      {"name": "Wellness", "name_id": "Kesehatan"},
		"educational":   {"name": "Educational", "name_id": "Edukasi"},
		"adventure":     {"name": "Adventure", "name_id": "Petualangan"},
	}

	categories := make([]fiber.Map, len(results))
	for i, result := range results {
		names, exists := categoryNames[result.Category]
		name := result.Category
		nameID := result.Category

		if exists {
			name = names["name"]
			nameID = names["name_id"]
		}

		categories[i] = fiber.Map{
			"id":      result.Category,
			"name":    name,
			"name_id": nameID,
			"count":   result.Count,
		}
	}

	return categories
}
