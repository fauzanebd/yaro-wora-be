package handlers

import (
	"fmt"
	"strconv"
	"time"
	"yaro-wora-be/config"
	"yaro-wora-be/models"
	"yaro-wora-be/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// =============================================================================
// ANALYTICS - ADMIN
// =============================================================================

// GetStorageAnalytics returns storage usage analytics
func GetStorageAnalytics(c *fiber.Ctx) error {
	// Get storage analytics
	analytics, err := utils.Storage.GetStorageAnalytics()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to get storage analytics: " + err.Error(),
			"code":    "INTERNAL_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    analytics,
	})
}

// TrackVisitor tracks a new visitor visit
func TrackVisitor(c *fiber.Ctx) error {
	// Get visitor information from request
	ipAddress := c.IP()
	userAgent := c.Get("User-Agent")
	referer := c.Get("Referer")
	page := c.Query("page", "/")
	sessionID := c.Get("X-Session-ID")

	// If no session ID provided, generate one
	if sessionID == "" {
		sessionID = utils.GenerateSessionID()
	}

	// Check if this is a unique visitor (first visit from this IP in the last 24 hours)
	var existingVisitor models.Visitor
	oneDayAgo := time.Now().Add(-24 * time.Hour)

	result := config.DB.Where("ip_address = ? AND created_at > ?", ipAddress, oneDayAgo).First(&existingVisitor)
	isUnique := result.Error == gorm.ErrRecordNotFound

	// Parse user agent for device/browser info (basic parsing)
	device := utils.ParseDevice(userAgent)
	browser := utils.ParseBrowser(userAgent)
	os := utils.ParseOS(userAgent)

	// Create visitor record
	visitor := models.Visitor{
		IPAddress: ipAddress,
		UserAgent: userAgent,
		Referer:   referer,
		Page:      page,
		SessionID: sessionID,
		Device:    device,
		Browser:   browser,
		OS:        os,
		IsUnique:  isUnique,
	}

	if err := config.DB.Create(&visitor).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to track visitor: " + err.Error(),
			"code":    "TRACKING_ERROR",
		})
	}

	return c.JSON(fiber.Map{
		"success":    true,
		"session_id": sessionID,
		"is_unique":  isUnique,
		"message":    "Visitor tracked successfully",
	})
}

// GetVisitorAnalytics returns comprehensive visitor analytics
func GetVisitorAnalytics(c *fiber.Ctx) error {
	db := config.DB

	// Get time range from query parameters (default: last 30 days)
	days := 30
	if daysStr := c.Query("days"); daysStr != "" {
		if parsedDays, err := strconv.Atoi(daysStr); err == nil && parsedDays > 0 {
			days = parsedDays
		}
	}

	startDate := time.Now().AddDate(0, 0, -days)
	today := time.Now().Truncate(24 * time.Hour)
	thisWeek := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))
	thisMonth := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -time.Now().Day()+1)

	// Total visitors
	var totalVisitors int64
	db.Model(&models.Visitor{}).Where("created_at >= ?", startDate).Count(&totalVisitors)

	// Unique visitors
	var uniqueVisitors int64
	db.Model(&models.Visitor{}).Where("created_at >= ? AND is_unique = ?", startDate, true).Count(&uniqueVisitors)

	// Today's visitors
	var todayVisitors int64
	db.Model(&models.Visitor{}).Where("created_at >= ?", today).Count(&todayVisitors)

	// This week's visitors
	var thisWeekVisitors int64
	db.Model(&models.Visitor{}).Where("created_at >= ?", thisWeek).Count(&thisWeekVisitors)

	// This month's visitors
	var thisMonthVisitors int64
	db.Model(&models.Visitor{}).Where("created_at >= ?", thisMonth).Count(&thisMonthVisitors)

	// Top pages
	var topPages []models.PageViewCount
	db.Model(&models.Visitor{}).
		Select("page, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("page").
		Order("count DESC").
		Limit(10).
		Scan(&topPages)

	// Top countries
	var topCountries []models.CountryViewCount
	db.Model(&models.Visitor{}).
		Select("country, COUNT(*) as count").
		Where("created_at >= ? AND country != ''", startDate).
		Group("country").
		Order("count DESC").
		Limit(10).
		Scan(&topCountries)

	// Top devices
	var topDevices []models.DeviceViewCount
	db.Model(&models.Visitor{}).
		Select("device, COUNT(*) as count").
		Where("created_at >= ? AND device != ''", startDate).
		Group("device").
		Order("count DESC").
		Limit(10).
		Scan(&topDevices)

	// Top browsers
	var topBrowsers []models.BrowserViewCount
	db.Model(&models.Visitor{}).
		Select("browser, COUNT(*) as count").
		Where("created_at >= ? AND browser != ''", startDate).
		Group("browser").
		Order("count DESC").
		Limit(10).
		Scan(&topBrowsers)

	// Hourly stats for today
	var hourlyStats []models.HourlyVisitorCount
	db.Model(&models.Visitor{}).
		Select("EXTRACT(hour FROM created_at) as hour, COUNT(*) as count").
		Where("created_at >= ?", today).
		Group("EXTRACT(hour FROM created_at)").
		Order("hour").
		Scan(&hourlyStats)

	// Daily stats
	var dailyStats []models.DailyVisitorCount
	db.Model(&models.Visitor{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE(created_at)").
		Order("date").
		Scan(&dailyStats)

	// Weekly stats
	var weeklyStats []models.WeeklyVisitorCount
	db.Model(&models.Visitor{}).
		Select("DATE_TRUNC('week', created_at) as week, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE_TRUNC('week', created_at)").
		Order("week").
		Scan(&weeklyStats)

	// Monthly stats
	var monthlyStats []models.MonthlyVisitorCount
	db.Model(&models.Visitor{}).
		Select("DATE_TRUNC('month', created_at) as month, COUNT(*) as count").
		Where("created_at >= ?", startDate).
		Group("DATE_TRUNC('month', created_at)").
		Order("month").
		Scan(&monthlyStats)

	analytics := models.VisitorAnalytics{
		TotalVisitors:     totalVisitors,
		UniqueVisitors:    uniqueVisitors,
		TodayVisitors:     todayVisitors,
		ThisWeekVisitors:  thisWeekVisitors,
		ThisMonthVisitors: thisMonthVisitors,
		TopPages:          topPages,
		TopCountries:      topCountries,
		TopDevices:        topDevices,
		TopBrowsers:       topBrowsers,
		HourlyStats:       hourlyStats,
		DailyStats:        dailyStats,
		WeeklyStats:       weeklyStats,
		MonthlyStats:      monthlyStats,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    analytics,
		"period":  fmt.Sprintf("Last %d days", days),
	})
}
