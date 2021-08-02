package models

import (
	"gorm.io/gorm"
	"time"
)

type BlogPost struct {
	ID        int `json:"id" gorm:"primarykey"`
	Title       string `json:"title"`
	Content  string `json:"content"`
	Slug      string `json:"slug"`
	Created_by string `json:"created_by"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}