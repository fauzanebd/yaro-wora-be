package models

import (
	"time"

	"gorm.io/gorm"
)

// Visitor represents a website visitor tracking record
type Visitor struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	IPAddress string         `json:"ip_address" gorm:"size:45;not null;index"`
	UserAgent string         `json:"user_agent" gorm:"type:text"`
	Referer   string         `json:"referer" gorm:"size:500"`
	Page      string         `json:"page" gorm:"size:200;index"`
	Country   string         `json:"country" gorm:"size:100"`
	City      string         `json:"city" gorm:"size:100"`
	Device    string         `json:"device" gorm:"size:50"` // mobile, desktop, tablet
	Browser   string         `json:"browser" gorm:"size:50"`
	OS        string         `json:"os" gorm:"size:50"`
	IsUnique  bool           `json:"is_unique" gorm:"default:true;index"`
	SessionID string         `json:"session_id" gorm:"size:100;index"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// VisitorAnalytics represents aggregated visitor analytics data
type VisitorAnalytics struct {
	TotalVisitors     int64                 `json:"total_visitors"`
	UniqueVisitors    int64                 `json:"unique_visitors"`
	TodayVisitors     int64                 `json:"today_visitors"`
	ThisWeekVisitors  int64                 `json:"this_week_visitors"`
	ThisMonthVisitors int64                 `json:"this_month_visitors"`
	TopPages          []PageViewCount       `json:"top_pages"`
	TopCountries      []CountryViewCount    `json:"top_countries"`
	TopDevices        []DeviceViewCount     `json:"top_devices"`
	TopBrowsers       []BrowserViewCount    `json:"top_browsers"`
	HourlyStats       []HourlyVisitorCount  `json:"hourly_stats"`
	DailyStats        []DailyVisitorCount   `json:"daily_stats"`
	WeeklyStats       []WeeklyVisitorCount  `json:"weekly_stats"`
	MonthlyStats      []MonthlyVisitorCount `json:"monthly_stats"`
}

// PageViewCount represents page view statistics
type PageViewCount struct {
	Page  string `json:"page"`
	Count int64  `json:"count"`
}

// CountryViewCount represents country view statistics
type CountryViewCount struct {
	Country string `json:"country"`
	Count   int64  `json:"count"`
}

// DeviceViewCount represents device view statistics
type DeviceViewCount struct {
	Device string `json:"device"`
	Count  int64  `json:"count"`
}

// BrowserViewCount represents browser view statistics
type BrowserViewCount struct {
	Browser string `json:"browser"`
	Count   int64  `json:"count"`
}

// HourlyVisitorCount represents hourly visitor statistics
type HourlyVisitorCount struct {
	Hour  int   `json:"hour"`
	Count int64 `json:"count"`
}

// DailyVisitorCount represents daily visitor statistics
type DailyVisitorCount struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// WeeklyVisitorCount represents weekly visitor statistics
type WeeklyVisitorCount struct {
	Week  string `json:"week"`
	Count int64  `json:"count"`
}

// MonthlyVisitorCount represents monthly visitor statistics
type MonthlyVisitorCount struct {
	Month string `json:"month"`
	Count int64  `json:"count"`
}

// VisitorSession represents a visitor session for tracking
type VisitorSession struct {
	SessionID string    `json:"session_id"`
	IPAddress string    `json:"ip_address"`
	StartTime time.Time `json:"start_time"`
	LastSeen  time.Time `json:"last_seen"`
	PageViews int       `json:"page_views"`
	IsActive  bool      `json:"is_active"`
}
