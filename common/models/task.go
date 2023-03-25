package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model         // adds ID, created_at etc.
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Comments    string `json:"comments"`
}
