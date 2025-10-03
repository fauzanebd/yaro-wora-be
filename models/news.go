package models

import (
	"time"

	"gorm.io/datatypes"
)

type NewsCategory struct {
	BaseModel
	Key           string `json:"key" gorm:"primaryKey;type:citext"`
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
	Color         string `json:"color" gorm:"default:#6b7280"`
	Icon          string `json:"icon"`
	SortOrder     int    `json:"sort_order" gorm:"default:0"`
}

type NewsArticle struct {
	BaseModel
	Title         string         `json:"title" gorm:"type:citext;not null"`
	TitleID       string         `json:"title_id"`
	Excerpt       string         `json:"excerpt" gorm:"type:text"`
	ExcerptID     string         `json:"excerpt_id" gorm:"type:text"`
	Content       string         `json:"content" gorm:"type:text;not null"`
	ContentID     string         `json:"content_id" gorm:"type:text"`
	AuthorName    string         `json:"author_name" gorm:"type:citext"`
	AuthorAvatar  string         `json:"author_avatar"`
	AuthorBio     string         `json:"author_bio"`
	AuthorEmail   string         `json:"author_email" gorm:"type:citext"`
	AuthorSocial  datatypes.JSON `json:"author_social" gorm:"type:jsonb"` // social links object
	DatePublished time.Time      `json:"date_published"`
	CategoryKey   string         `json:"category_key" gorm:"type:citext"`
	Category      NewsCategory   `json:"category" gorm:"foreignKey:CategoryKey;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FeaturedImage string         `json:"featured_image"`
	ImageGallery  datatypes.JSON `json:"image_gallery" gorm:"type:jsonb"` // array of Image objects
	Tags          datatypes.JSON `json:"tags" gorm:"type:jsonb"`          // array of strings
	ReadTime      int            `json:"read_time" gorm:"default:5"`      // in minutes
	IsHeadline    bool           `json:"is_headline" gorm:"default:false"`
	ViewCount     int            `json:"view_count" gorm:"default:0"`
	Language      string         `json:"language" gorm:"default:en"`
	SEOMetaTitle  string         `json:"seo_meta_title"`
	SEOMetaDesc   string         `json:"seo_meta_desc"`
	SEOKeywords   datatypes.JSON `json:"seo_keywords" gorm:"type:jsonb"` // array of strings
	CanonicalURL  string         `json:"canonical_url"`
	SortOrder     int            `json:"sort_order" gorm:"default:0"`
}

// Override the BaseModel ID field for string ID
func (NewsCategory) TableName() string {
	return "news_categories"
}

func (NewsArticle) TableName() string {
	return "news_articles"
}
