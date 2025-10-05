package models

import "gorm.io/datatypes"

type Attraction struct {
	BaseModel
	Title              string         `json:"title" gorm:"type:citext;not null"`
	TitleID            string         `json:"title_id" gorm:"type:citext;not null"`
	ShortDescription   string         `json:"short_description" gorm:"type:citext"`
	ShortDescriptionID string         `json:"short_description_id" gorm:"type:citext"`
	FullDescription    string         `json:"full_description" gorm:"type:citext"`
	FullDescriptionID  string         `json:"full_description_id" gorm:"type:citext"`
	ImageURL           string         `json:"image_url"`
	Highlights         datatypes.JSON `json:"highlights" gorm:"type:jsonb"`    // array of strings
	HighlightsID       datatypes.JSON `json:"highlights_id" gorm:"type:jsonb"` // array of strings
	DurationMinutes    int            `json:"duration_minutes"`
	Difficulty         string         `json:"difficulty"` // easy, medium, hard
	DifficultyID       string         `json:"difficulty_id" gorm:"type:citext"`
	PriceRange         string         `json:"price_range"` // "50000-100000"
	IsFeatured         bool           `json:"is_featured" gorm:"default:false"`
	SortOrder          int            `json:"sort_order" gorm:"default:0"`
}

// Override the BaseModel ID field for string ID
func (Attraction) TableName() string {
	return "attractions"
}
