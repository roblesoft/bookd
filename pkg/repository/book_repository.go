package repository

import (
	"fmt"

	"github.com/roblesoft/bookd/pkg/models"
)

type BookRepository struct {
	GormRepository
}

func (r *BookRepository) List(args map[string]any) (any, error) {
	var wc models.BookCollection
	order := "created_at"
	err := r.db.Limit(args["limit"].(int)).Order(order).Where(fmt.Sprintf("%v > ?", order), args["after"]).Limit(args["limit"].(int)).Find(&wc).Error

	return wc, err
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

func (r *BookRepository) Update(args any) (any, error) {
	w := args.(map[string]any)["entity"].(*models.Book)

	if err := r.db.Model(w).Where("id = ?", args.(map[string]any)["id"]).Updates(w).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *BookRepository) Delete(id any) (bool, error) {

	garden, err := r.Get(id)

	if err != nil {
		return false, err
	}

	if err := r.db.Delete(garden).Error; err != nil {
		return false, err
	}

	return true, nil
}
