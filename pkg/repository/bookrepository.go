package repository

import (
	"strconv"

	"github.com/roblesoft/bookd/pkg/models"
)

type BookRepository struct {
	GormRepository
}

func (r *BookRepository) GetAll(limit, offset string) (any, error) {
	var books []models.Book
	offsetInt, _ := strconv.Atoi(offset)
	limitInt, _ := strconv.Atoi(limit)
	err := r.db.Find(&books).Error
	if err := r.db.Limit(limitInt).Offset(offsetInt).Find(&books).Error; err != nil {
		return nil, err
	}
	return books, err
}

func (r *BookRepository) Get(id any) (any, error) {
	var w *models.Book

	err := r.db.Where("id = ?", id).First(&w).Error

	return w, err
}

func (r *BookRepository) Create(book *models.Book) error {
	err := r.db.Create(book).Error
	return err
}

func (r *BookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *BookRepository) Delete(id any) error {
	err := r.db.Delete(&models.Book{}, id).Error
	return err
}
