package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `gorm:"size:255;not null"`
	Slug        string `gorm:"size:255;unique_index;not null;"`
	Description string `gorm:"size:255;not null"`
	Duration    int
	Image       string `gorm:"size:255;not null"`
}
