package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	Author    string         `json:"author"`
	Year      int            `json:"year"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index,->" json:"-"`
}

type BookCollection []*Book
