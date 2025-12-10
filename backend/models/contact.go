package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Name    string `json:"name" gorm:"not null"`
	Email   string `json:"email" gorm:"not null"`
	Subject string `json:"subject"`
	Message string `json:"message" gorm:"type:text;not null"`
	Status  string `json:"status" gorm:"default:'new'"` // new, read, responded
}
