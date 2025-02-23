package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Answer   string `json:"answer"`
	UserID   uint   `json:"user_id"`
	IssueURL string `json:"issue_url"`
}
