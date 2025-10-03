package models

import "gorm.io/datatypes"

type GalleryCategory struct {
	BaseModel
	ID            string `json:"id" gorm:"primaryKey;type:citext"`
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
	Color         string `json:"color" gorm:"default:#6b7280"`
	Icon          string `json:"icon"`
	SortOrder     int    `json:"sort_order" gorm:"default:0"`
}

type GalleryImage struct {
	BaseModel
	ID           string          `json:"id" gorm:"primaryKey;type:citext"`
	Title        string          `json:"title" gorm:"type:citext;not null"`
	Description  string          `json:"description" gorm:"type:text"`
	ImageURL     string          `json:"image_url" gorm:"not null"`
	ThumbnailURL string          `json:"thumbnail_url"`
	HighResURL   string          `json:"high_res_url"`
	CategoryID   string          `json:"category_id" gorm:"type:citext"`
	Category     GalleryCategory `json:"category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Dimensions   datatypes.JSON  `json:"dimensions" gorm:"type:jsonb"` // {width, height, aspect_ratio}
	Photographer string          `json:"photographer" gorm:"type:citext"`
	Location     string          `json:"location" gorm:"type:citext"`
	Tags         datatypes.JSON  `json:"tags" gorm:"type:jsonb"`     // array of strings
	Metadata     datatypes.JSON  `json:"metadata" gorm:"type:jsonb"` // camera info, settings, etc
	SortOrder    int             `json:"sort_order" gorm:"default:0"`
}

// Override the BaseModel ID field for string ID
func (GalleryCategory) TableName() string {
	return "gallery_categories"
}

func (GalleryImage) TableName() string {
	return "gallery_images"
}
