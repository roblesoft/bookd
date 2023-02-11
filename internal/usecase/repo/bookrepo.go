package repository

import (
	"strconv"

	entity "github.com/roblesoft/bookd/internal/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	Db *gorm.DB
}

func (r *BookRepository) GetAll(limit, offset string) (any, error) {
	var books []entity.Book
	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)
	err := r.Db.Find(&books).Error
	if err := r.Db.Limit(limitInt).Offset(offsetInt).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, err
}

func (r *BookRepository) Get(id any) (any, error) {
	var w *entity.Book

	err := r.Db.Where("id = ?", id).First(&w).Error

	return w, err
}

func (r *BookRepository) Create(book *entity.Book) error {
	err := r.Db.Create(book).Error
	return err
}

func (r *BookRepository) Update(book *entity.Book) error {
	return r.Db.Save(book).Error
}

func (r *BookRepository) Delete(id any) error {
	err := r.Db.Delete(&entity.Book{}, id).Error
	return err
}
