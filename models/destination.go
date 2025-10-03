package models

import "gorm.io/datatypes"

type Destination struct {
	BaseModel
	ID                  string         `json:"id" gorm:"primaryKey;type:citext"`
	Title               string         `json:"title" gorm:"type:citext;not null"`
	TitleID             string         `json:"title_id"`
	Description         string         `json:"description" gorm:"type:text"`
	DescriptionID       string         `json:"description_id" gorm:"type:text"`
	FullContent         string         `json:"full_content" gorm:"type:text"`
	FullContentID       string         `json:"full_content_id" gorm:"type:text"`
	ImageURL            string         `json:"image_url"`
	ThumbnailURL        string         `json:"thumbnail_url"`
	HeroImageURL        string         `json:"hero_image_url"`
	Category            string         `json:"category" gorm:"type:citext"` // nature, culture, heritage, agriculture, adventure
	CategoryID          string         `json:"category_id"`
	Type                string         `json:"type"` // main, other
	LocationLat         float64        `json:"location_lat"`
	LocationLng         float64        `json:"location_lng"`
	LocationAddress     string         `json:"location_address"`
	LocationAddressID   string         `json:"location_address_id"`
	GoogleMapsURL       string         `json:"google_maps_url"`
	Highlights          datatypes.JSON `json:"highlights" gorm:"type:jsonb"`    // array of strings
	HighlightsID        datatypes.JSON `json:"highlights_id" gorm:"type:jsonb"` // array of strings
	Facilities          datatypes.JSON `json:"facilities" gorm:"type:jsonb"`    // array of strings
	FacilitiesID        datatypes.JSON `json:"facilities_id" gorm:"type:jsonb"` // array of strings
	Images              datatypes.JSON `json:"images" gorm:"type:jsonb"`        // array of Image objects
	BestVisitTime       string         `json:"best_visit_time"`                 // morning, afternoon, evening
	BestVisitTimeDesc   string         `json:"best_visit_time_desc"`
	BestVisitTimeDescID string         `json:"best_visit_time_desc_id"`
	Duration            string         `json:"duration"`
	DurationDesc        string         `json:"duration_desc"`
	DurationDescID      string         `json:"duration_desc_id"`
	Difficulty          string         `json:"difficulty"`    // easy, medium, hard
	Accessibility       string         `json:"accessibility"` // wheelchair_friendly, etc
	EntranceFee         *int           `json:"entrance_fee"`
	OperatingHours      datatypes.JSON `json:"operating_hours" gorm:"type:jsonb"` // OperatingHours object
	BookingRequired     bool           `json:"booking_required" gorm:"default:false"`
	ContactInfo         datatypes.JSON `json:"contact_info" gorm:"type:jsonb"` // SimpleContactInfo object
	IsFeatured          bool           `json:"is_featured" gorm:"default:false"`
	SortOrder           int            `json:"sort_order" gorm:"default:0"`
}

// Override the BaseModel ID field for string ID
func (Destination) TableName() string {
	return "destinations"
}
