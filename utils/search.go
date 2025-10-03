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
