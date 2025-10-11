package models

import "gorm.io/datatypes"

type ProfilePageContent struct {
	BaseModel
	Title                  string         `json:"title" gorm:"not null"`
	TitleID                string         `json:"title_id"`
	HeaderImageURL         string         `json:"header_image_url"`
	Subtitle               string         `json:"subtitle" gorm:"type:text"`
	SubtitleID             string         `json:"subtitle_id"`
	BriefSectionTitle      string         `json:"brief_section_title"`
	BriefSectionTitleID    string         `json:"brief_section_title_id"`
	BriefSectionContent    string         `json:"brief_section_content" gorm:"type:text"`
	BriefSectionContentID  string         `json:"brief_section_content_id"`
	BriefSectionImageURL   string         `json:"brief_section_image_url"`
	ProfileSections        datatypes.JSON `json:"profile_sections" gorm:"type:jsonb"` // array of ProfileSection objects
	CTASectionTitle        string         `json:"cta_section_title"`
	CTASectionTitleID      string         `json:"cta_section_title_id"`
	CTASectionText         string         `json:"cta_section_text"`
	CTASectionTextID       string         `json:"cta_section_text_id"`
	CTASectionButtonText   string         `json:"cta_section_button_text"`
	CTASectionButtonTextID string         `json:"cta_section_button_text_id"`
	CTASectionButtonURL    string         `json:"cta_section_button_url"`
}

type ProfileSection struct {
	Title     string `json:"title"`
	TitleID   string `json:"title_id"`
	Content   string `json:"content" gorm:"type:text"` // will be rendered as markdown
	ContentID string `json:"content_id"`
	ImageURL  string `json:"image_url"`
}
