package models

type Pricing struct {
	BaseModel
	Type               string `json:"type" gorm:"unique;not null"` // domestic, locals_sumba, foreigner
	Title              string `json:"title" gorm:"not null"`
	TitleID            string `json:"title_id" gorm:"not null"`
	Subtitle           string `json:"subtitle"`
	SubtitleID         string `json:"subtitle_id"`
	AdultPrice         int    `json:"adult_price" gorm:"not null"`
	InfantPrice        int    `json:"infant_price" gorm:"not null"`
	Currency           string `json:"currency" gorm:"default:IDR"`
	Description        string `json:"description"`
	ImageURL           string `json:"image_url"`
	ThumbnailURL       string `json:"thumbnail_url"`
	PrimaryColor       string `json:"color"`
	StartGradientColor string `json:"start_gradient_color"`
	EndGradientColor   string `json:"end_gradient_color"`
}

type GeneralPricingContent struct {
	BaseModel
	GeneralPricingSectionTitle         string `json:"general_pricing_section_title"`
	GeneralPricingSectionTitleID       string `json:"general_pricing_section_title_id"`
	GeneralPricingSectionDescription   string `json:"general_pricing_section_description"`
	GeneralPricingSectionDescriptionID string `json:"general_pricing_section_description_id"`
}
