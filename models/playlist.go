package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model

	Link        string `json:"link"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image_link"`
	TodoID      uint
}
