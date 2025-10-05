package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type RegulationCategory struct {
	Key           string         `json:"key" gorm:"primaryKey;type:citext"`
	Name          string         `json:"name" gorm:"type:citext;not null"`
	NameID        string         `json:"name_id"`
	Description   string         `json:"description"`
	DescriptionID string         `json:"description_id"`
	Icon          string         `json:"icon"`
	Color         string         `json:"color" gorm:"default:#6b7280"`
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type Regulation struct {
	BaseModel
	ID           string             `json:"id" gorm:"primaryKey;type:text"`
	CategoryKey  string             `json:"category_key" gorm:"type:citext"`
	Category     RegulationCategory `json:"category" gorm:"foreignKey:CategoryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Question     string             `json:"question" gorm:"not null"`
	Answer       string             `json:"answer" gorm:"type:text;not null"`
	SearchVector string             `json:"-" gorm:"type:tsvector;index:,type:gin"`
	Priority     int                `json:"priority" gorm:"default:0"`
	IsActive     bool               `json:"is_active" gorm:"default:true"`
	Tags         datatypes.JSON     `json:"tags" gorm:"type:jsonb"` // array of strings
}

func (RegulationCategory) TableName() string {
	return "regulation_categories"
}

func (Regulation) TableName() string {
	return "regulations"
}

// BeforeCreate hook to update search vector
func (r *Regulation) BeforeCreate(tx *gorm.DB) error {
	return r.updateSearchVector(tx)
}

// BeforeUpdate hook to update search vector
func (r *Regulation) BeforeUpdate(tx *gorm.DB) error {
	return r.updateSearchVector(tx)
}

func (r *Regulation) updateSearchVector(tx *gorm.DB) error {
	sql := `
		UPDATE regulations 
		SET search_vector = 
			setweight(to_tsvector('english', COALESCE(question, '')), 'A') ||
			setweight(to_tsvector('english', COALESCE(answer, '')), 'B')
		WHERE id = ?
	`
	return tx.Exec(sql, r.ID).Error
}
