package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Email       string `json:"email"`
	Password       string `json:"password"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}