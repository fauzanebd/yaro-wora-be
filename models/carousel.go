package models

type Carousel struct {
	BaseModel
	Title         string `json:"title" gorm:"not null"`
	TitleID       string `json:"title_id" gorm:"not null"`
	Subtitle      string `json:"subtitle"`
	SubtitleID    string `json:"subtitle_id"`
	ImageURL      string `json:"image_url" gorm:"not null"`
	ThumbnailURL  string `json:"thumbnail_url"`
	AltText       string `json:"alt_text"`
	AltTextID     string `json:"alt_text_id"`
	CarouselOrder int    `json:"carousel_order" gorm:"default:0"`
	IsActive      bool   `json:"is_active" gorm:"default:true"`
}
