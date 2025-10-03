package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel contains common fields for all models
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Image represents an image with metadata
type Image struct {
	URL     string `json:"url"`
	Caption string `json:"caption,omitempty"`
	AltText string `json:"alt_text,omitempty"`
}

// Location represents geographical coordinates
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Address   string  `json:"address"`
	AddressID string  `json:"address_id,omitempty"`
}

// OperatingHours represents business hours
type OperatingHours struct {
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}

// SimpleContactInfo represents basic contact information
type SimpleContactInfo struct {
	Phone    string `json:"phone"`
	WhatsApp string `json:"whatsapp"`
	Email    string `json:"email"`
}
