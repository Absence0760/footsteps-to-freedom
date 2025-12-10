package models

import (
	"time"

	"gorm.io/gorm"
)

// Item represents a generic item in the database
// Customize this model based on your project needs
type Item struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name" gorm:"not null"`
	Description string       `json:"description"`
	Status    string         `json:"status" gorm:"default:'active'"`
}

// TableName overrides the table name
func (Item) TableName() string {
	return "items"
}
