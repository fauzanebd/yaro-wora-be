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
	WhyVisitSectionTitlePart1    string `json:"why_visit_section_title_part_1"`
	WhyVisitSectionTitlePart1ID  string `json:"why_visit_section_title_part_1_id"`
	WhyVisitSectionTitlePart2    string `json:"why_visit_section_title_part_2"`
	WhyVisitSectionTitlePart2ID  string `json:"why_visit_section_title_part_2_id"`
	WhyVisitSectionDescription   string `json:"why_visit_section_description"`
	WhyVisitSectionDescriptionID string `json:"why_visit_section_description_id"`
}
