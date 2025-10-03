package models

import "gorm.io/datatypes"

type Attraction struct {
	BaseModel
	ID               string         `json:"id" gorm:"primaryKey;type:citext"`
	Title            string         `json:"title" gorm:"type:citext;not null"`
	ShortDescription string         `json:"short_description"`
	FullDescription  string         `json:"full_description" gorm:"type:text"`
	ImageURL         string         `json:"image_url"`
	Highlights       datatypes.JSON `json:"highlights" gorm:"type:jsonb"` // array of strings
	Duration         string         `json:"duration"`
	Difficulty       string         `json:"difficulty"`  // easy, medium, hard
	PriceRange       string         `json:"price_range"` // "50000-100000"
	IsFeatured       bool           `json:"is_featured" gorm:"default:false"`
	SortOrder        int            `json:"sort_order" gorm:"default:0"`
}

// Override the BaseModel ID field for string ID
func (Attraction) TableName() string {
	return "attractions"
}
