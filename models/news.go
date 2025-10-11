package models

import (
	"errors"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type NewsCategory struct {
	BaseModel
	Name          string `json:"name" gorm:"type:citext;not null"`
	NameID        string `json:"name_id"`
	Description   string `json:"description"`
	DescriptionID string `json:"description_id"`
}

type NewsAuthor struct {
	BaseModel
	Name   string `json:"name" gorm:"type:citext;not null"`
	Avatar string `json:"avatar"`
}

type NewsArticle struct {
	BaseModel
	Title          string         `json:"title" gorm:"type:citext;not null"`
	TitleID        string         `json:"title_id"`
	Excerpt        string         `json:"excerpt" gorm:"type:text"`
	ExcerptID      string         `json:"excerpt_id" gorm:"type:text"`
	Content        string         `json:"content" gorm:"type:text;not null"` // markdown content
	ContentID      string         `json:"content_id" gorm:"type:text"`       // markdown content
	SearchVectorEN string         `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_en"`
	SearchVectorID string         `json:"-" gorm:"type:tsvector;index:,type:gin;column:search_vector_id"`
	AuthorID       uint           `json:"author_id"`
	NewsAuthor     NewsAuthor     `json:"news_author" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DatePublished  time.Time      `json:"date_published"`
	CategoryID     uint           `json:"category_id"`
	NewsCategory   NewsCategory   `json:"news_category" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ImageURL       string         `json:"image_url"`
	Tags           datatypes.JSON `json:"tags" gorm:"type:jsonb"`     // array of strings
	ReadTime       int            `json:"read_time" gorm:"default:5"` // in minutes
	IsHeadline     bool           `json:"is_headline" gorm:"default:false"`
}

type NewsArticleSummary struct {
	ID            uint           `json:"id"`
	Title         string         `json:"title"`
	TitleID       string         `json:"title_id"`
	Excerpt       string         `json:"excerpt"`
	ExcerptID     string         `json:"excerpt_id"`
	AuthorID      uint           `json:"author_id"`
	NewsAuthor    NewsAuthor     `json:"news_author"`
	DatePublished time.Time      `json:"date_published"`
	CategoryID    uint           `json:"category_id"`
	NewsCategory  NewsCategory   `json:"news_category"`
	ImageURL      string         `json:"image_url"`
	Tags          datatypes.JSON `json:"tags"`
	ReadTime      int            `json:"read_time"`
	IsHeadline    bool           `json:"is_headline"`
}

type NewsPageContent struct {
	BaseModel
	HeroImageURL            string `json:"hero_image_url"`
	HeroImageThumbnailURL   string `json:"hero_image_thumbnail_url"`
	Title                   string `json:"title"`
	TitleID                 string `json:"title_id"`
	Subtitle                string `json:"subtitle"`
	SubtitleID              string `json:"subtitle_id"`
	HighlightSectionTitle   string `json:"highlight_section_title"`
	HighlightSectionTitleID string `json:"highlight_section_title_id"`
}

func (NewsArticle) TableName() string {
	return "news_articles"
}

// BeforeCreate hook to update search vector
func (n *NewsArticle) BeforeCreate(tx *gorm.DB) error {
	if err := n.ensureSingleHighlighted(tx, true); err != nil {
		return err
	}
	return n.updateSearchVector(tx)
}

// BeforeUpdate hook to update search vector
func (n *NewsArticle) BeforeUpdate(tx *gorm.DB) error {
	if err := n.ensureSingleHighlighted(tx, false); err != nil {
		return err
	}
	return n.updateSearchVector(tx)
}

// ensureSingleHighlighted validates that only one destination can have is_headline = true
func (d *NewsArticle) ensureSingleHighlighted(tx *gorm.DB, isCreate bool) error {
	if !d.IsHeadline {
		return nil
	}

	var count int64
	query := tx.Model(&NewsArticle{}).Where("is_headline = ?", true)
	if !isCreate && d.ID != 0 {
		query = query.Where("id <> ?", d.ID)
	}
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("only one news article can be highlighted at a time")
	}
	return nil
}

func (n *NewsArticle) updateSearchVector(tx *gorm.DB) error {
	// Update both English and Indonesian search vectors
	sql := `
		UPDATE news_articles 
		SET 
			search_vector_en = 
				setweight(to_tsvector('english', COALESCE(news_articles.title, '')), 'A') ||
				setweight(to_tsvector('english', COALESCE(news_articles.excerpt, '')), 'B') ||
				setweight(to_tsvector('english', COALESCE(news_articles.content, '')), 'C') ||
				setweight(to_tsvector('simple', COALESCE(author.name, '')), 'D'),
			search_vector_id = 
				setweight(to_tsvector('simple', COALESCE(news_articles.title_id, '')), 'A') ||
				setweight(to_tsvector('simple', COALESCE(news_articles.excerpt_id, '')), 'B') ||
				setweight(to_tsvector('simple', COALESCE(news_articles.content_id, '')), 'C') ||
				setweight(to_tsvector('simple', COALESCE(author.name, '')), 'D')
		FROM news_authors author
		WHERE news_articles.author_id = author.id AND news_articles.id = ?
	`
	return tx.Exec(sql, n.ID).Error
}
