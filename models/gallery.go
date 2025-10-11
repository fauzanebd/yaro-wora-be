package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GalleryCategory struct {
	BaseModel
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
}

type GalleryImage struct {
	BaseModel
	Title              string          `json:"title" gorm:"type:citext;not null"`
	TitleID            string          `json:"title_id"`
	ShortDescription   string          `json:"short_description" gorm:"type:text"`
	ShortDescriptionID string          `json:"short_description_id" gorm:"type:text"`
	Description        string          `json:"description" gorm:"type:text"`
	DescriptionID      string          `json:"description_id" gorm:"type:text"`
	SearchVectorEN     string          `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_en"`
	SearchVectorID     string          `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_id"`
	ImageURL           string          `json:"image_url"`
	ThumbnailURL       string          `json:"thumbnail_url"`
	CategoryID         uint            `json:"category_id"`
	GalleryCategory    GalleryCategory `json:"gallery_category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Photographer       string          `json:"photographer" gorm:"type:citext"`
	Location           string          `json:"location" gorm:"type:citext"`
	Tags               datatypes.JSON  `json:"tags" gorm:"type:jsonb"`    // array of strings
	TagsID             datatypes.JSON  `json:"tags_id" gorm:"type:jsonb"` // array of strings
	DateUploaded       time.Time       `json:"date_uploaded"`
}

type GalleryImageSummary struct {
	ID                 uint            `json:"id"`
	Title              string          `json:"title" gorm:"type:citext;not null"`
	TitleID            string          `json:"title_id"`
	ShortDescription   string          `json:"short_description" gorm:"type:text"`
	ShortDescriptionID string          `json:"short_description_id" gorm:"type:text"`
	ImageURL           string          `json:"image_url"`
	ThumbnailURL       string          `json:"thumbnail_url"`
	CategoryID         uint            `json:"category_id"`
	GalleryCategory    GalleryCategory `json:"gallery_category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DateUploaded       time.Time       `json:"date_uploaded"`
}

type GalleryPageContent struct {
	BaseModel
	HeroImageURL          string `json:"hero_image_url"`
	HeroImageThumbnailURL string `json:"hero_image_thumbnail_url"`
	Title                 string `json:"title"`
	TitleID               string `json:"title_id"`
	Subtitle              string `json:"subtitle"`
	SubtitleID            string `json:"subtitle_id"`
}

// Override the BaseModel ID field for string ID
func (GalleryCategory) TableName() string {
	return "gallery_categories"
}

func (GalleryImage) TableName() string {
	return "gallery_images"
}

// BeforeCreate hook to update search vector
func (d *GalleryImage) BeforeCreate(tx *gorm.DB) error {
	return d.updateSearchVector(tx)
}

// BeforeUpdate hook to update search vector
func (d *GalleryImage) BeforeUpdate(tx *gorm.DB) error {
	return d.updateSearchVector(tx)
}

func (d *GalleryImage) updateSearchVector(tx *gorm.DB) error {
	sql := `
		UPDATE gallery_images 
		SET 
			search_vector_en = 
				setweight(to_tsvector('english', COALESCE(title, '')), 'A') ||
				setweight(to_tsvector('english', COALESCE(short_description, '')), 'B') ||
				setweight(to_tsvector('english', COALESCE(description, '')), 'C'),
			search_vector_id = 
				setweight(to_tsvector('simple', COALESCE(title_id, '')), 'A') ||
				setweight(to_tsvector('simple', COALESCE(short_description_id, '')), 'B') ||
				setweight(to_tsvector('simple', COALESCE(description_id, '')), 'C')
		WHERE id = ?
	`
	return tx.Exec(sql, d.ID).Error
}
