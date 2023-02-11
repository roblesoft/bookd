package usecase

import (
	entity "github.com/roblesoft/bookd/internal/entity"
	repo "github.com/roblesoft/bookd/internal/usecase/repo"
)

type Service struct {
	bookRepo *repo.BookRepository
}

func NewService(bookRepo *repo.BookRepository) *Service {
	return &Service{
		bookRepo: bookRepo,
	}
}

func (s *Service) GetAllBooks(limit, offset string) ([]entity.Book, error) {
	return s.bookRepo.GetAll(limit, offset)
}

func (s *Service) GetBook(id int) (*entity.Book, error) {
	return s.bookRepo.Get(id)
}

func (s *Service) CreateBook(b *entity.Book) error {
	return s.bookRepo.Create(b)
}

func (s *Service) UpdateBook(b *entity.Book) error {
	return s.bookRepo.Update(b)
}

func (s *Service) DeleteBook(id int) error {
	return s.bookRepo.Delete(id)
}
