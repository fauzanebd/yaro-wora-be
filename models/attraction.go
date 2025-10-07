package models

import "gorm.io/datatypes"

type Attraction struct {
	BaseModel
	Title         string         `json:"title" gorm:"type:citext;not null"`
	TitleID       string         `json:"title_id" gorm:"type:citext;not null"`
	Subtitle      string         `json:"subtitle" gorm:"type:citext"`
	SubtitleID    string         `json:"subtitle_id" gorm:"type:citext"`
	Description   string         `json:"description" gorm:"type:citext"`
	DescriptionID string         `json:"description_id" gorm:"type:citext"`
	ImageURL      string         `json:"image_url"`
	Highlights    datatypes.JSON `json:"highlights" gorm:"type:jsonb"`    // array of strings
	HighlightsID  datatypes.JSON `json:"highlights_id" gorm:"type:jsonb"` // array of strings
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
	Active        bool           `json:"active" gorm:"default:true"`
}

type GeneralAttractionContent struct {
	BaseModel
	AttractionSectionTitlePart1    string `json:"attraction_section_title_part_1"`
	AttractionSectionTitlePart1ID  string `json:"attraction_section_title_part_1_id"`
	AttractionSectionTitlePart2    string `json:"attraction_section_title_part_2"`
	AttractionSectionTitlePart2ID  string `json:"attraction_section_title_part_2_id"`
	AttractionSectionDescription   string `json:"attraction_section_description"`
	AttractionSectionDescriptionID string `json:"attraction_section_description_id"`
}

// Override the BaseModel ID field for string ID
func (Attraction) TableName() string {
	return "attractions"
}
