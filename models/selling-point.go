package models

type SellingPoint struct {
	BaseModel
	Title             string `json:"title" gorm:"not null"`
	TitleID           string `json:"title_id" gorm:"not null"`
	Description       string `json:"description"`
	DescriptionID     string `json:"description_id"`
	ImageURL          string `json:"image_url" gorm:"not null"`
	ThumbnailURL      string `json:"thumbnail_url"`
	PillarColor       string `json:"pillar_color"`
	TextColor         string `json:"text_color"`
	SellingPointOrder int    `json:"selling_point_order" gorm:"default:0"`
	IsActive          bool   `json:"is_active" gorm:"default:true"`
}
