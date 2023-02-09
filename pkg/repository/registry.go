package repository

import (
	"gorm.io/gorm"
)

type Repository interface {
	Configure(*gorm.DB)
	List(map[string]any) (any, error)
	Get(any) (any, error)
	Create(any) (any, error)
	Update(any) (any, error)
	Delete(any) (bool, error)
}

type GormRepository struct {
	db *gorm.DB
}

func (r *GormRepository) Configure(db *gorm.DB) {
	r.db = db
}
