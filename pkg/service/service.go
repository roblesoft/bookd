package repo

import "github.com/roblesoft/bookd/pkg/repository"

type Service struct {
	bookRepo *repository.BookRepository
}

func NewService(bookRepo *repository.BookRepository) *Service {
	return &Service{
		bookRepo: bookRepo,
	}
}
