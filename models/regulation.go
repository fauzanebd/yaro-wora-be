package models

import (
	"gorm.io/gorm"
)

type RegulationPageContent struct {
	BaseModel
	HeroImageURL          string `json:"hero_image_url"`
	HeroImageThumbnailURL string `json:"hero_image_thumbnail_url"`
	Title                 string `json:"title"`
	TitleID               string `json:"title_id"`
	Subtitle              string `json:"subtitle"`
	SubtitleID            string `json:"subtitle_id"`
	CTATitle              string `json:"cta_title"`
	CTATitleID            string `json:"cta_title_id"`
	CTADescription        string `json:"cta_description"`
	CTADescriptionID      string `json:"cta_description_id"`
	CTAButtonText         string `json:"cta_button_text"`
	CTAButtonTextID       string `json:"cta_button_text_id"`
	CTAButtonURL          string `json:"cta_button_url"`
}

type RegulationCategory struct {
	BaseModel
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
}

type Regulation struct {
	BaseModel
	CategoryID         uint               `json:"category_id"`
	RegulationCategory RegulationCategory `json:"regulation_category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question           string             `json:"question" gorm:"not null"`
	QuestionID         string             `json:"question_id" gorm:"type:text"`
	Answer             string             `json:"answer" gorm:"type:text;not null"`
	AnswerID           string             `json:"answer_id" gorm:"type:text"`
	SearchVectorEN     string             `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_en"`
	SearchVectorID     string             `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_id"`
}

func (RegulationCategory) TableName() string {
	return "regulation_categories"
}

func (Regulation) TableName() string {
	return "regulations"
}

// BeforeCreate hook to update search vector
func (r *Regulation) BeforeCreate(tx *gorm.DB) error {
	return r.updateSearchVector(tx)
}

// BeforeUpdate hook to update search vector
func (r *Regulation) BeforeUpdate(tx *gorm.DB) error {
	return r.updateSearchVector(tx)
}

func (r *Regulation) updateSearchVector(tx *gorm.DB) error {
	sql := `
		UPDATE regulations 
		SET 
			search_vector_en = 
				setweight(to_tsvector('english', COALESCE(question, '')), 'A') ||
				setweight(to_tsvector('english', COALESCE(answer, '')), 'B'),
			search_vector_id = 
				setweight(to_tsvector('simple', COALESCE(question_id, '')), 'A') ||
				setweight(to_tsvector('simple', COALESCE(answer_id, '')), 'B')
		WHERE id = ?
	`
	return tx.Exec(sql, r.ID).Error
}
