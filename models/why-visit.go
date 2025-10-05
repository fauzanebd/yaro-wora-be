package models

type WhyVisit struct {
	BaseModel
	Title         string `json:"title" gorm:"not null"`
	TitleID       string `json:"title_id" gorm:"not null"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
	IconURL       string `json:"icon_url" gorm:"not null"`
}

type GeneralWhyVisitContent struct {
	BaseModel
	WhyVisitSectionTitle         string `json:"why_visit_section_title"`
	WhyVisitSectionTitleID       string `json:"why_visit_section_title_id"`
	WhyVisitSectionDescription   string `json:"why_visit_section_description"`
	WhyVisitSectionDescriptionID string `json:"why_visit_section_description_id"`
}
