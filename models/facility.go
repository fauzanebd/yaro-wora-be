package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type FacilityCategory struct {
	BaseModel
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description" gorm:"type:text"`
	DescriptionID string `json:"description_id"`
}

type Facility struct {
	BaseModel
	Name                   string           `json:"name" gorm:"type:citext;not null"`
	NameID                 string           `json:"name_id"`
	ShortDescription       string           `json:"short_description" gorm:"type:text"`
	ShortDescriptionID     string           `json:"short_description_id" gorm:"type:text"`
	Description            string           `json:"description" gorm:"type:text"`
	DescriptionID          string           `json:"description_id" gorm:"type:text"`
	SearchVectorEN         string           `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_en"`
	SearchVectorID         string           `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_id"`
	ImageURL               string           `json:"image_url"`
	ThumbnailURL           string           `json:"thumbnail_url"`
	CategoryID             uint             `json:"category_id"`
	FacilityCategory       FacilityCategory `json:"facility_category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FacilityDetailSections datatypes.JSON   `json:"facility_detail_sections" gorm:"type:jsonb"`
	Highlights             datatypes.JSON   `json:"highlights" gorm:"type:jsonb"`    // array of strings
	HighlightsID           datatypes.JSON   `json:"highlights_id" gorm:"type:jsonb"` // array of strings
	Duration               string           `json:"duration"`
	Capacity               string           `json:"capacity"`
	Price                  string           `json:"price"`
	DurationID             string           `json:"duration_id"`
	CapacityID             string           `json:"capacity_id"`
	PriceID                string           `json:"price_id"`
	CTAUrl                 string           `json:"cta_url"`
	IsFeatured             bool             `json:"is_featured" gorm:"default:false"`
	SortOrder              int              `json:"sort_order" gorm:"default:0"`
}

type FacilitySummary struct {
	ID                 uint             `json:"id"`
	Name               string           `json:"name"`
	NameID             string           `json:"name_id"`
	ShortDescription   string           `json:"short_description"`
	ShortDescriptionID string           `json:"short_description_id"`
	ImageURL           string           `json:"image_url"`
	ThumbnailURL       string           `json:"thumbnail_url"`
	Highlights         datatypes.JSON   `json:"highlights"`
	HighlightsID       datatypes.JSON   `json:"highlights_id"`
	IsFeatured         bool             `json:"is_featured"`
	SortOrder          int              `json:"sort_order"`
	CategoryID         uint             `json:"category_id"`
	FacilityCategory   FacilityCategory `json:"facility_category"`
	Duration           string           `json:"duration"`
	Capacity           string           `json:"capacity"`
	Price              string           `json:"price"`
	DurationID         string           `json:"duration_id"`
	CapacityID         string           `json:"capacity_id"`
	PriceID            string           `json:"price_id"`
}

type FacilityDetailSection struct {
	Title     string `json:"title"`
	TitleID   string `json:"title_id"`
	Content   string `json:"content"`
	ContentID string `json:"content_id"`
	ImageURL  string `json:"image_url"`
}

type FacilityPageContent struct {
	BaseModel
	HeroImageURL                       string `json:"hero_image_url"`
	HeroImageThumbnailURL              string `json:"hero_image_thumbnail_url"`
	Title                              string `json:"title"`
	TitleID                            string `json:"title_id"`
	Subtitle                           string `json:"subtitle"`
	SubtitleID                         string `json:"subtitle_id"`
	FacilitiesListSectionTitle         string `json:"facilities_list_section_title"`
	FacilitiesListSectionTitleID       string `json:"facilities_list_section_title_id"`
	FacilitiesListSectionDescription   string `json:"facilities_list_section_description"`
	FacilitiesListSectionDescriptionID string `json:"facilities_list_section_description_id"`
	CTATitle                           string `json:"cta_title"`
	CTATitleID                         string `json:"cta_title_id"`
	CTADescription                     string `json:"cta_description"`
	CTADescriptionID                   string `json:"cta_description_id"`
	CTAButtonText                      string `json:"cta_button_text"`
	CTAButtonTextID                    string `json:"cta_button_text_id"`
	CTAButtonURL                       string `json:"cta_button_url"`
}

func (Facility) TableName() string {
	return "facilities"
}

// BeforeCreate hook to update search vector
func (f *Facility) BeforeCreate(tx *gorm.DB) error {
	return f.updateSearchVector(tx)
}

// BeforeUpdate hook to update search vector
func (f *Facility) BeforeUpdate(tx *gorm.DB) error {
	return f.updateSearchVector(tx)
}

func (f *Facility) updateSearchVector(tx *gorm.DB) error {
	sql := `
		UPDATE facilities 
		SET 
			search_vector_en = 
				setweight(to_tsvector('english', COALESCE(name, '')), 'A') ||
				setweight(to_tsvector('english', COALESCE(description, '')), 'B') ||
				setweight(
					to_tsvector(
						'english',
						COALESCE(
							(
								SELECT string_agg(
									(elem->>'title') || ' ' || (elem->>'content'),
									' '
								)
								FROM jsonb_array_elements(facility_detail_sections) AS elem
							),
							''
						)
					),
					'C'
				),
			search_vector_id = 
				setweight(to_tsvector('simple', COALESCE(name_id, '')), 'A') ||
				setweight(to_tsvector('simple', COALESCE(description_id, '')), 'B') ||
				setweight(
					to_tsvector(
						'simple',
						COALESCE(
							(
								SELECT string_agg(
									(elem->>'title_id') || ' ' || (elem->>'content_id'),
									' '
								)
								FROM jsonb_array_elements(facility_detail_sections) AS elem
							),
							''
						)
					),
					'C'
				)
		WHERE id = ?
	`
	return tx.Exec(sql, f.ID).Error
}
