package models

import "gorm.io/datatypes"

// type ContactSubmission struct {
// 	BaseModel
// 	Name          string         `json:"name" gorm:"not null"`
// 	Email         string         `json:"email" gorm:"type:citext;not null"`
// 	Phone         string         `json:"phone"`
// 	Subject       string         `json:"subject" gorm:"not null"`
// 	Message       string         `json:"message" gorm:"type:text;not null"`
// 	PreferredDate string         `json:"preferred_date"`
// 	VisitorType   string         `json:"visitor_type"`                    // domestic, locals_sumba, foreigner
// 	VisitorCount  datatypes.JSON `json:"visitor_count" gorm:"type:jsonb"` // {adults: int, infants: int}
// 	Status        string         `json:"status" gorm:"default:pending"`   // pending, reviewed, responded, closed
// 	ReferenceID   string         `json:"reference_id" gorm:"unique"`
// 	AdminNotes    string         `json:"admin_notes" gorm:"type:text"`
// 	ResponseSent  bool           `json:"response_sent" gorm:"default:false"`
// }

type ContactInfo struct {
	BaseModel
	AddressPart1     string         `json:"address_part_1"`
	AddressPart1ID   string         `json:"address_part_1_id"`
	AddressPart2     string         `json:"address_part_2"`
	AddressPart2ID   string         `json:"address_part_2_id"`
	Latitude         float64        `json:"latitude"`
	Longitude        float64        `json:"longitude"`
	Phones           datatypes.JSON `json:"phones" gorm:"type:jsonb"`       // array of strings
	Emails           datatypes.JSON `json:"emails" gorm:"type:jsonb"`       // array of strings
	SocialMedia      datatypes.JSON `json:"social_media" gorm:"type:jsonb"` // SocialMedia object
	PlanYourVisitURL string         `json:"plan_your_visit_url"`
}

type SocialMedia struct {
	Name    string `json:"name"`
	Handle  string `json:"handle"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type ContactContent struct {
	BaseModel
	ContactSectionTitlePart1    string `json:"contact_section_title_part_1"`
	ContactSectionTitlePart1ID  string `json:"contact_section_title_part_1_id"`
	ContactSectionTitlePart2    string `json:"contact_section_title_part_2"`
	ContactSectionTitlePart2ID  string `json:"contact_section_title_part_2_id"`
	ContactSectionDescription   string `json:"contact_section_description"`
	ContactSectionDescriptionID string `json:"contact_section_description_id"`
}
