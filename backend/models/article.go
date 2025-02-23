package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"userId"`
	GitPath string `json:"gitPath"`
}
