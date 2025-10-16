package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model

	Title       string   `json:"title"`
	Description string   `json:"description"`
	Completed   bool     `json:"completed"`
	Playlist    Playlist `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"playlist"`
}
