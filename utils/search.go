package utils

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// SearchHelper provides utility functions for database searches
type SearchHelper struct{}

var Search = &SearchHelper{}

// CaseInsensitiveSearch performs case-insensitive search using ILIKE for non-citext fields
// For citext fields, regular equality (=) works case-insensitively
func (s *SearchHelper) CaseInsensitiveSearch(db *gorm.DB, field, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s ILIKE ?", field), "%"+value+"%")
}

// ExactCaseInsensitiveMatch for exact matches (useful for emails, usernames with citext)
// With citext fields, this automatically works case-insensitively
func (s *SearchHelper) ExactCaseInsensitiveMatch(db *gorm.DB, field, value string) *gorm.DB {
	return db.Where(fmt.Sprintf("%s = ?", field), value)
}

// MultiFieldSearch searches across multiple fields with case-insensitive ILIKE
func (s *SearchHelper) MultiFieldSearch(db *gorm.DB, searchTerm string, fields []string) *gorm.DB {
	if searchTerm == "" || len(fields) == 0 {
		return db
	}

	// Build OR conditions for multiple fields
	conditions := make([]string, len(fields))
	args := make([]interface{}, len(fields))

	for i, field := range fields {
		conditions[i] = fmt.Sprintf("%s ILIKE ?", field)
		args[i] = "%" + searchTerm + "%"
	}

	query := strings.Join(conditions, " OR ")
	return db.Where(query, args...)
}

// EmailSearch optimized for citext email fields
func (s *SearchHelper) EmailSearch(db *gorm.DB, email string) *gorm.DB {
	// With citext, this automatically works case-insensitively
	return db.Where("email = ?", email)
}

// UsernameSearch optimized for citext username fields
func (s *SearchHelper) UsernameSearch(db *gorm.DB, username string) *gorm.DB {
	// With citext, this automatically works case-insensitively
	return db.Where("username = ?", username)
}

// SearchConfig represents search configuration
type SearchConfig struct {
	Query         string   `json:"query"`
	Fields        []string `json:"fields"`
	ExactMatch    bool     `json:"exact_match"`
	CaseSensitive bool     `json:"case_sensitive"`
}

// AdvancedSearch performs configurable search operations
func (s *SearchHelper) AdvancedSearch(db *gorm.DB, config SearchConfig) *gorm.DB {
	if config.Query == "" {
		return db
	}

	if config.ExactMatch {
		// For exact matches
		if len(config.Fields) == 1 {
			if config.CaseSensitive {
				return db.Where(fmt.Sprintf("%s = ?", config.Fields[0]), config.Query)
			} else {
				// Use ILIKE for case-insensitive exact match on non-citext fields
				// For citext fields, regular = already works case-insensitively
				return db.Where(fmt.Sprintf("%s ILIKE ?", config.Fields[0]), config.Query)
			}
		}
		return db
	}

	// For partial matches
	return s.MultiFieldSearch(db, config.Query, config.Fields)
}

// ValidateEmail checks if email format is valid
func (s *SearchHelper) ValidateEmail(email string) bool {
	// Basic email validation
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// NormalizeSearchTerm normalizes search terms
func (s *SearchHelper) NormalizeSearchTerm(term string) string {
	// Trim whitespace and convert to lowercase for consistency
	return strings.TrimSpace(strings.ToLower(term))
}

// GallerySearch optimized for citext gallery fields
func (s *SearchHelper) GallerySearch(db *gorm.DB, searchTerm string) *gorm.DB {
	if searchTerm == "" {
		return db
	}

	// Search across multiple citext fields in gallery
	fields := []string{"title", "photographer", "location"}
	return s.MultiFieldSearch(db, searchTerm, fields)
}

// CategorySearch for case-insensitive category matching
func (s *SearchHelper) CategorySearch(db *gorm.DB, category string) *gorm.DB {
	// With citext category fields, this works case-insensitively
	return db.Where("category = ?", category)
}

// IDSearch for case-insensitive ID matching (useful for string primary keys)
func (s *SearchHelper) IDSearch(db *gorm.DB, id string) *gorm.DB {
	// With citext ID fields, this works case-insensitively
	return db.Where("id = ?", id)
}

// FullTextSearch performs PostgreSQL full-text search using tsvector
// This is MUCH faster than ILIKE for large text content
// Deprecated: Use FullTextSearchLang instead for multi-language support
func (s *SearchHelper) FullTextSearch(db *gorm.DB, searchTerm string) *gorm.DB {
	return s.FullTextSearchLang(db, searchTerm, "en")
}

// FullTextSearchLang performs full-text search with language support
// lang: "en" for English, "id" for Indonesian
func (s *SearchHelper) FullTextSearchLang(db *gorm.DB, searchTerm string, lang string) *gorm.DB {
	if searchTerm == "" {
		return db
	}

	// Convert search term to tsquery format
	// Multiple words are ANDed together with &
	query := strings.ReplaceAll(strings.TrimSpace(searchTerm), " ", " & ")

	// Choose the appropriate search vector column based on language
	vectorColumn := "search_vector_en"
	tsConfig := "english"
	if lang == "id" || lang == "ID" {
		vectorColumn = "search_vector_id"
		tsConfig = "simple" // Indonesian uses simple config (no stemming)
	}

	return db.Where(fmt.Sprintf("%s @@ to_tsquery('%s', ?)", vectorColumn, tsConfig), query)
}

// FullTextSearchWithRanking performs full-text search and orders by relevance
// Deprecated: Use FullTextSearchWithRankingLang instead for multi-language support
func (s *SearchHelper) FullTextSearchWithRanking(db *gorm.DB, searchTerm string) *gorm.DB {
	return s.FullTextSearchWithRankingLang(db, searchTerm, "en")
}

// FullTextSearchWithRankingLang performs full-text search with ranking and language support
// lang: "en" for English, "id" for Indonesian
func (s *SearchHelper) FullTextSearchWithRankingLang(db *gorm.DB, searchTerm string, lang string) *gorm.DB {
	if searchTerm == "" {
		return db
	}

	query := strings.ReplaceAll(strings.TrimSpace(searchTerm), " ", " & ")

	// Choose the appropriate search vector column based on language
	vectorColumn := "search_vector_en"
	tsConfig := "english"
	if lang == "id" || lang == "ID" {
		vectorColumn = "search_vector_id"
		tsConfig = "simple"
	}

	whereClause := fmt.Sprintf("%s @@ to_tsquery('%s', ?)", vectorColumn, tsConfig)
	rankExpr := fmt.Sprintf("ts_rank(%s, to_tsquery('%s', ?))", vectorColumn, tsConfig)

	return db.
		Where(whereClause, query).
		Order(gorm.Expr(rankExpr+" DESC", query))
}

// FullTextSearchOr performs full-text search with OR logic (any word matches)
// Deprecated: Use FullTextSearchOrLang instead for multi-language support
func (s *SearchHelper) FullTextSearchOr(db *gorm.DB, searchTerm string) *gorm.DB {
	return s.FullTextSearchOrLang(db, searchTerm, "en")
}

// FullTextSearchOrLang performs full-text search with OR logic and language support
// lang: "en" for English, "id" for Indonesian
func (s *SearchHelper) FullTextSearchOrLang(db *gorm.DB, searchTerm string, lang string) *gorm.DB {
	if searchTerm == "" {
		return db
	}

	// Multiple words are ORed together with |
	query := strings.ReplaceAll(strings.TrimSpace(searchTerm), " ", " | ")

	// Choose the appropriate search vector column based on language
	vectorColumn := "search_vector_en"
	tsConfig := "english"
	if lang == "id" || lang == "ID" {
		vectorColumn = "search_vector_id"
		tsConfig = "simple"
	}

	return db.Where(fmt.Sprintf("%s @@ to_tsquery('%s', ?)", vectorColumn, tsConfig), query)
}

// UpdateSearchVector updates the search_vector field for full-text search
// Call this after creating or updating records
func (s *SearchHelper) UpdateSearchVector(db *gorm.DB, tableName string, fields []string, id interface{}) error {
	// Build tsvector expression combining multiple fields with weights
	var vectorParts []string
	for i, field := range fields {
		// Weight: A (highest) for first field, B for second, C for third, D for rest
		weight := "D"
		switch i {
		case 0:
			weight = "A"
		case 1:
			weight = "B"
		case 2:
			weight = "C"
		}
		vectorParts = append(vectorParts, fmt.Sprintf("setweight(to_tsvector('english', COALESCE(%s, '')), '%s')", field, weight))
	}

	vectorExpr := strings.Join(vectorParts, " || ")

	sql := fmt.Sprintf("UPDATE %s SET search_vector = %s WHERE id = ?", tableName, vectorExpr)
	return db.Exec(sql, id).Error
}

// UpdateNewsSearchVector updates search vector for news articles
func (s *SearchHelper) UpdateNewsSearchVector(db *gorm.DB, id interface{}) error {
	fields := []string{"title", "excerpt", "content"}
	return s.UpdateSearchVector(db, "news_articles", fields, id)
}

// UpdateDestinationSearchVector updates search vector for destinations
func (s *SearchHelper) UpdateDestinationSearchVector(db *gorm.DB, id interface{}) error {
	fields := []string{"title", "description", "full_content"}
	return s.UpdateSearchVector(db, "destinations", fields, id)
}

// UpdateFacilitySearchVector updates search vector for facilities
func (s *SearchHelper) UpdateFacilitySearchVector(db *gorm.DB, id interface{}) error {
	fields := []string{"name", "description", "full_content"}
	return s.UpdateSearchVector(db, "facilities", fields, id)
}

// UpdateRegulationSearchVector updates search vector for regulations
func (s *SearchHelper) UpdateRegulationSearchVector(db *gorm.DB, id interface{}) error {
	fields := []string{"question", "answer"}
	return s.UpdateSearchVector(db, "regulations", fields, id)
}
