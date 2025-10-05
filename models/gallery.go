package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GalleryCategory struct {
	ID            string         `json:"id" gorm:"primaryKey;type:text"`
	Name          string         `json:"name" gorm:"type:citext;not null"`
	NameID        string         `json:"name_id"`
	Description   string         `json:"description"`
	DescriptionID string         `json:"description_id"`
	Color         string         `json:"color" gorm:"default:#6b7280"`
	Icon          string         `json:"icon"`
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type GalleryImage struct {
	ID           string          `json:"id" gorm:"primaryKey;type:text"`
	Title        string          `json:"title" gorm:"type:citext;not null"`
	Description  string          `json:"description" gorm:"type:text"`
	ImageURL     string          `json:"image_url" gorm:"not null"`
	ThumbnailURL string          `json:"thumbnail_url"`
	HighResURL   string          `json:"high_res_url"`
	CategoryID   string          `json:"category_id" gorm:"type:text"`
	Category     GalleryCategory `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dimensions   datatypes.JSON  `json:"dimensions" gorm:"type:jsonb"` // {width, height, aspect_ratio}
	Photographer string          `json:"photographer" gorm:"type:citext"`
	Location     string          `json:"location" gorm:"type:citext"`
	Tags         datatypes.JSON  `json:"tags" gorm:"type:jsonb"`     // array of strings
	Metadata     datatypes.JSON  `json:"metadata" gorm:"type:jsonb"` // camera info, settings, etc
	SortOrder    int             `json:"sort_order" gorm:"default:0"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"-" gorm:"index"`
}

// Override the BaseModel ID field for string ID
func (GalleryCategory) TableName() string {
	return "gallery_categories"
}

func (GalleryImage) TableName() string {
	return "gallery_images"
}
