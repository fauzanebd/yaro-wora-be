package models

import (
	"gorm.io/datatypes"
)

type Heritage struct {
	BaseModel
	Title                  string         `json:"title" gorm:"type:citext;not null"`
	TitleID                string         `json:"title_id"`
	ShortDescription       string         `json:"short_description" gorm:"type:text"`
	ShortDescriptionID     string         `json:"short_description_id" gorm:"type:text"`
	Description            string         `json:"description" gorm:"type:text"`
	DescriptionID          string         `json:"description_id" gorm:"type:text"`
	ImageURL               string         `json:"image_url"`
	ThumbnailURL           string         `json:"thumbnail_url"`
	HeritageDetailSections datatypes.JSON `json:"heritage_detail_sections" gorm:"type:jsonb"`
	SortOrder              int            `json:"sort_order" gorm:"default:0"`
}

type HeritageSummary struct {
	ID                 uint   `json:"id"`
	Title              string `json:"title"`
	TitleID            string `json:"title_id"`
	ShortDescription   string `json:"short_description"`
	ShortDescriptionID string `json:"short_description_id"`
	ImageURL           string `json:"image_url"`
	ThumbnailURL       string `json:"thumbnail_url"`
	SortOrder          int    `json:"sort_order"`
}

type HeritageDetailSection struct {
	Title     string `json:"title"`
	TitleID   string `json:"title_id"`
	Content   string `json:"content"`
	ContentID string `json:"content_id"`
	ImageURL  string `json:"image_url"`
}

type HeritagePageContent struct {
	BaseModel
	HeroImageURL             string `json:"hero_image_url"`
	HeroImageThumbnailURL    string `json:"hero_image_thumbnail_url"`
	Title                    string `json:"title"`
	TitleID                  string `json:"title_id"`
	Subtitle                 string `json:"subtitle"`
	SubtitleID               string `json:"subtitle_id"`
	MainSectionTitle         string `json:"main_section_title"`
	MainSectionTitleID       string `json:"main_section_title_id"`
	MainSectionDescription   string `json:"main_section_description"`
	MainSectionDescriptionID string `json:"main_section_description_id"`
	CTATitle                 string `json:"cta_title"`
	CTATitleID               string `json:"cta_title_id"`
	CTADescription           string `json:"cta_description"`
	CTADescriptionID         string `json:"cta_description_id"`
	CTAButtonText            string `json:"cta_button_text"`
	CTAButtonTextID          string `json:"cta_button_text_id"`
	CTAButtonURL             string `json:"cta_button_url"`
}

func (Heritage) TableName() string {
	return "heritages"
}
