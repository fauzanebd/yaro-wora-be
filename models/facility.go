package models

import "gorm.io/datatypes"

type Facility struct {
	BaseModel
	ID                 string         `json:"id" gorm:"primaryKey;type:citext"`
	Name               string         `json:"name" gorm:"type:citext;not null"`
	NameID             string         `json:"name_id"`
	Description        string         `json:"description" gorm:"type:text"`
	DescriptionID      string         `json:"description_id" gorm:"type:text"`
	FullContent        string         `json:"full_content" gorm:"type:text"`
	FullContentID      string         `json:"full_content_id" gorm:"type:text"`
	ImageURL           string         `json:"image_url"`
	ThumbnailURL       string         `json:"thumbnail_url"`
	Category           string         `json:"category" gorm:"type:citext"` // accommodation, workshop, culinary, entertainment, activity, wellness, educational, adventure
	CategoryID         string         `json:"category_id"`
	Highlights         datatypes.JSON `json:"highlights" gorm:"type:jsonb"`    // array of strings
	HighlightsID       datatypes.JSON `json:"highlights_id" gorm:"type:jsonb"` // array of strings
	Duration           string         `json:"duration"`
	Capacity           string         `json:"capacity"`
	Price              int            `json:"price"`
	Currency           string         `json:"currency" gorm:"default:IDR"`
	BookingRequired    bool           `json:"booking_required" gorm:"default:false"`
	AdvanceBookingDays int            `json:"advance_booking_days" gorm:"default:0"`
	Availability       datatypes.JSON `json:"availability" gorm:"type:jsonb"`     // availability object
	Includes           datatypes.JSON `json:"includes" gorm:"type:jsonb"`         // array of strings
	IncludesID         datatypes.JSON `json:"includes_id" gorm:"type:jsonb"`      // array of strings
	Requirements       datatypes.JSON `json:"requirements" gorm:"type:jsonb"`     // array of strings
	RequirementsID     datatypes.JSON `json:"requirements_id" gorm:"type:jsonb"`  // array of strings
	WhatToBring        datatypes.JSON `json:"what_to_bring" gorm:"type:jsonb"`    // array of strings
	WhatToBringID      datatypes.JSON `json:"what_to_bring_id" gorm:"type:jsonb"` // array of strings
	Images             datatypes.JSON `json:"images" gorm:"type:jsonb"`           // array of Image objects
	LocationLat        float64        `json:"location_lat"`
	LocationLng        float64        `json:"location_lng"`
	LocationAddress    string         `json:"location_address"`
	LocationAddressID  string         `json:"location_address_id"`
	BookingPolicy      datatypes.JSON `json:"booking_policy" gorm:"type:jsonb"` // booking policy object
	ContactInfo        datatypes.JSON `json:"contact_info" gorm:"type:jsonb"`   // SimpleContactInfo object
	SortOrder          int            `json:"sort_order" gorm:"default:0"`
}

type Booking struct {
	BaseModel
	BookingID             string   `json:"booking_id" gorm:"unique;not null"`
	FacilityID            string   `json:"facility_id"`
	Facility              Facility `json:"facility" gorm:"foreignKey:FacilityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GuestName             string   `json:"guest_name" gorm:"not null"`
	Email                 string   `json:"email" gorm:"type:citext;not null"`
	Phone                 string   `json:"phone" gorm:"not null"`
	CheckInDate           string   `json:"check_in_date"`
	CheckOutDate          string   `json:"check_out_date"`
	Participants          int      `json:"participants" gorm:"default:1"`
	SpecialRequirements   string   `json:"special_requirements" gorm:"type:text"`
	LanguagePreference    string   `json:"language_preference" gorm:"default:en"`
	TotalPrice            int      `json:"total_price"`
	Currency              string   `json:"currency" gorm:"default:IDR"`
	Status                string   `json:"status" gorm:"default:pending"` // pending, confirmed, cancelled, completed
	ConfirmationEmailSent bool     `json:"confirmation_email_sent" gorm:"default:false"`
}

// Override the BaseModel ID field for string ID
func (Facility) TableName() string {
	return "facilities"
}
