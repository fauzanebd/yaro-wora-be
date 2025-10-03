package models

type Pricing struct {
	BaseModel
	Type        string `json:"type" gorm:"unique;not null"` // domestic, locals_sumba, foreigner
	Title       string `json:"title" gorm:"not null"`
	Subtitle    string `json:"subtitle"`
	AdultPrice  int    `json:"adult_price" gorm:"not null"`
	InfantPrice int    `json:"infant_price" gorm:"not null"`
	Currency    string `json:"currency" gorm:"default:IDR"`
	Description string `json:"description"`
}
