package models

import (
	common "api/lib"
	"github.com/jinzhu/gorm"
)

// Email data model
type Email struct {
	gorm.Model
	Email               string
	EmailHeader         string
	NumberOfOccurrences int
}

// Serialize serializes user data
func (e *Email) Serialize() common.JSON {
	return common.JSON{
		"id":                e.ID,
		"email":             e.Email,
		"email_header":      e.EmailHeader,
		"no_of_occurrences": e.NumberOfOccurrences,
		"date_created":      e.CreatedAt,
		"updated_at":        e.UpdatedAt,
	}
}

func (e *Email) Read(m common.JSON) {
	e.ID = uint(m["id"].(float64))
	e.Email = m["email"].(string)
	e.EmailHeader = m["email_header"].(string)
	e.NumberOfOccurrences = m["no_of_occurrences"].(int)
}
