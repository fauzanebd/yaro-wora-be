package models

import "gorm.io/datatypes"

type ContactSubmission struct {
	BaseModel
	Name          string         `json:"name" gorm:"not null"`
	Email         string         `json:"email" gorm:"type:citext;not null"`
	Phone         string         `json:"phone"`
	Subject       string         `json:"subject" gorm:"not null"`
	Message       string         `json:"message" gorm:"type:text;not null"`
	PreferredDate string         `json:"preferred_date"`
	VisitorType   string         `json:"visitor_type"`                    // domestic, locals_sumba, foreigner
	VisitorCount  datatypes.JSON `json:"visitor_count" gorm:"type:jsonb"` // {adults: int, infants: int}
	Status        string         `json:"status" gorm:"default:pending"`   // pending, reviewed, responded, closed
	ReferenceID   string         `json:"reference_id" gorm:"unique"`
	AdminNotes    string         `json:"admin_notes" gorm:"type:text"`
	ResponseSent  bool           `json:"response_sent" gorm:"default:false"`
}

type ContactInfo struct {
	BaseModel
	Street         string         `json:"street"`
	City           string         `json:"city"`
	Province       string         `json:"province"`
	Country        string         `json:"country"`
	PostalCode     string         `json:"postal_code"`
	Latitude       float64        `json:"latitude"`
	Longitude      float64        `json:"longitude"`
	Phones         datatypes.JSON `json:"phones" gorm:"type:jsonb"` // array of strings
	Emails         datatypes.JSON `json:"emails" gorm:"type:jsonb"` // array of strings
	WhatsApp       string         `json:"whatsapp"`
	SocialMedia    datatypes.JSON `json:"social_media" gorm:"type:jsonb"`    // social media object
	OperatingHours datatypes.JSON `json:"operating_hours" gorm:"type:jsonb"` // OperatingHours object
}
