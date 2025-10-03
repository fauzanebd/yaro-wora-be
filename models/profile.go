package models

import "gorm.io/datatypes"

type Profile struct {
	BaseModel
	Title             string         `json:"title" gorm:"not null"`
	Description       string         `json:"description" gorm:"type:text"`
	VisionTitle       string         `json:"vision_title"`
	VisionContent     string         `json:"vision_content" gorm:"type:text"`
	MissionTitle      string         `json:"mission_title"`
	MissionContent    string         `json:"mission_content" gorm:"type:text"`
	ObjectivesTitle   string         `json:"objectives_title"`
	ObjectivesContent string         `json:"objectives_content" gorm:"type:text"`
	FeaturedImages    datatypes.JSON `json:"featured_images" gorm:"type:jsonb"` // array of Image objects
}

// VisionMissionObjective represents the structured content
type VisionMissionObjective struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
