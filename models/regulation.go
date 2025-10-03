package models

import "gorm.io/datatypes"

type RegulationCategory struct {
	BaseModel
	Key           string `json:"key" gorm:"primaryKey;type:citext"`
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
	Icon          string `json:"icon"`
	Color         string `json:"color" gorm:"default:#6b7280"`
	SortOrder     int    `json:"sort_order" gorm:"default:0"`
}

type Regulation struct {
	BaseModel
	ID          string             `json:"id" gorm:"primaryKey;type:citext"`
	CategoryKey string             `json:"category_key" gorm:"type:citext"`
	Category    RegulationCategory `json:"category" gorm:"foreignKey:CategoryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question    string             `json:"question" gorm:"not null"`
	Answer      string             `json:"answer" gorm:"type:text;not null"`
	Priority    int                `json:"priority" gorm:"default:0"`
	IsActive    bool               `json:"is_active" gorm:"default:true"`
	Tags        datatypes.JSON     `json:"tags" gorm:"type:jsonb"` // array of strings
}

// Override the BaseModel ID field for string ID
func (RegulationCategory) TableName() string {
	return "regulation_categories"
}

func (Regulation) TableName() string {
	return "regulations"
}
