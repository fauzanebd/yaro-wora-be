package models

import (
	"errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type DestinationCategory struct {
	BaseModel
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
}

type Destination struct {
	BaseModel
	Title                     string              `json:"title" gorm:"type:citext;not null"`
	TitleID                   string              `json:"title_id"`
	ShortDescription          string              `json:"short_description" gorm:"type:text"`
	ShortDescriptionID        string              `json:"short_description_id" gorm:"type:text"`
	About                     string              `json:"about" gorm:"type:text"`
	AboutID                   string              `json:"about_id" gorm:"type:text"`
	SearchVectorEN            string              `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_en"`
	SearchVectorID            string              `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_id"`
	ImageURL                  string              `json:"image_url"`
	ThumbnailURL              string              `json:"thumbnail_url"`
	DestinationDetailSections datatypes.JSON      `json:"destination_detail_sections" gorm:"type:jsonb"`
	Highlights                datatypes.JSON      `json:"highlights" gorm:"type:jsonb"`
	HighlightsID              datatypes.JSON      `json:"highlights_id" gorm:"type:jsonb"`
	CTAUrl                    string              `json:"cta_url"`
	GoogleMapsURL             string              `json:"google_maps_url"`
	IsFeatured                bool                `json:"is_featured" gorm:"default:false"`
	SortOrder                 int                 `json:"sort_order" gorm:"default:0"`
	CategoryID                uint                `json:"category_id"`
	DestinationCategory       DestinationCategory `json:"destination_category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type DestinationSummary struct {
	ID                  uint                `json:"id"`
	Title               string              `json:"title"`
	TitleID             string              `json:"title_id"`
	ShortDescription    string              `json:"short_description"`
	ShortDescriptionID  string              `json:"short_description_id"`
	ImageURL            string              `json:"image_url"`
	ThumbnailURL        string              `json:"thumbnail_url"`
	Highlights          datatypes.JSON      `json:"highlights"`
	HighlightsID        datatypes.JSON      `json:"highlights_id"`
	IsFeatured          bool                `json:"is_featured"`
	SortOrder           int                 `json:"sort_order"`
	CategoryID          uint                `json:"category_id"`
	DestinationCategory DestinationCategory `json:"destination_category"`
}

type DestinationDetailSection struct {
	Title     string `json:"title"`
	TitleID   string `json:"title_id"`
	Content   string `json:"content"`
	ContentID string `json:"content_id"`
	ImageURL  string `json:"image_url"`
}

type DestinationPageContent struct {
	BaseModel
	HeroImageURL                     string `json:"hero_image_url"`
	HeroImageThumbnailURL            string `json:"hero_image_thumbnail_url"`
	Title                            string `json:"title"`
	TitleID                          string `json:"title_id"`
	Subtitle                         string `json:"subtitle"`
	SubtitleID                       string `json:"subtitle_id"`
	FeaturedDestinationTitle         string `json:"featured_destination_title"`
	FeaturedDestinationTitleID       string `json:"featured_destination_title_id"`
	FeaturedDestinationDescription   string `json:"featured_destination_description"`
	FeaturedDestinationDescriptionID string `json:"featured_destination_description_id"`
	OtherDestinationsTitle           string `json:"other_destinations_title"`
	OtherDestinationsTitleID         string `json:"other_destinations_title_id"`
	OtherDestinationsDescription     string `json:"other_destinations_description"`
	OtherDestinationsDescriptionID   string `json:"other_destinations_description_id"`
	CTATitle                         string `json:"cta_title"`
	CTATitleID                       string `json:"cta_title_id"`
	CTADescription                   string `json:"cta_description"`
	CTADescriptionID                 string `json:"cta_description_id"`
	CTAButtonText                    string `json:"cta_button_text"`
	CTAButtonTextID                  string `json:"cta_button_text_id"`
	CTAButtonURL                     string `json:"cta_button_url"`
}

func (Destination) TableName() string {
	return "destinations"
}

// BeforeCreate hook to update search vector
func (d *Destination) BeforeCreate(tx *gorm.DB) error {
	if err := d.ensureSingleFeatured(tx, true); err != nil {
		return err
	}
	return d.updateSearchVector(tx)
}

// BeforeUpdate hook to update search vector
func (d *Destination) BeforeUpdate(tx *gorm.DB) error {
	if err := d.ensureSingleFeatured(tx, false); err != nil {
		return err
	}
	return d.updateSearchVector(tx)
}

func (d *Destination) updateSearchVector(tx *gorm.DB) error {
	sql := `
		UPDATE destinations 
		SET 
			search_vector_en = 
				setweight(to_tsvector('english', COALESCE(title, '')), 'A') ||
				setweight(to_tsvector('english', COALESCE(about, '')), 'B') ||
				setweight(
					to_tsvector(
						'english',
						COALESCE(
							(
								SELECT string_agg(
									(elem->>'title') || ' ' || (elem->>'content'),
									' '
								)
								FROM jsonb_array_elements(destination_detail_sections) AS elem
							),
							''
						)
					),
					'C'
				),
			search_vector_id = 
				setweight(to_tsvector('simple', COALESCE(title_id, '')), 'A') ||
				setweight(to_tsvector('simple', COALESCE(about_id, '')), 'B') ||
				setweight(
					to_tsvector(
						'simple',
						COALESCE(
							(
								SELECT string_agg(
									(elem->>'title_id') || ' ' || (elem->>'content_id'),
									' '
								)
								FROM jsonb_array_elements(destination_detail_sections) AS elem
							),
							''
						)
					),
					'C'
				)
		WHERE id = ?
	`
	return tx.Exec(sql, d.ID).Error
}

// ensureSingleFeatured validates that only one destination can have is_featured = true
func (d *Destination) ensureSingleFeatured(tx *gorm.DB, isCreate bool) error {
	if !d.IsFeatured {
		return nil
	}

	var count int64
	query := tx.Model(&Destination{}).Where("is_featured = ?", true)
	if !isCreate && d.ID != 0 {
		query = query.Where("id <> ?", d.ID)
	}
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("only one destination can be featured at a time")
	}
	return nil
}
