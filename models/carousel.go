package models

type Carousel struct {
	BaseModel
	Title    string `json:"title" gorm:"not null"`
	Subtitle string `json:"subtitle"`
	ImageURL string `json:"image_url" gorm:"not null"`
	AltText  string `json:"alt_text"`
	Order    int    `json:"order" gorm:"default:0"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}
