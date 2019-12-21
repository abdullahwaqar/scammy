package model

import "github.com/jinzhu/gorm"

type Email struct {
	gorm.Model
	Email               string
	EmailHeader         string
	NumberOfOccurrences string
}
