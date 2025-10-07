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
	GeneralPricingSectionTitlePart1    string `json:"general_pricing_section_title_part_1"`
	GeneralPricingSectionTitlePart1ID  string `json:"general_pricing_section_title_part_1_id"`
	GeneralPricingSectionTitlePart2    string `json:"general_pricing_section_title_part_2"`
	GeneralPricingSectionTitlePart2ID  string `json:"general_pricing_section_title_part_2_id"`
	GeneralPricingSectionDescription   string `json:"general_pricing_section_description"`
	GeneralPricingSectionDescriptionID string `json:"general_pricing_section_description_id"`
}
