package models

import "gorm.io/gorm"

// Defines
type Todo struct {
	gorm.Model
	Name        string
	Description string
}
