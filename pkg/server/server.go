package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesoft/bookd/pkg/repository"
)

type Server struct {
	Port       string
	App        *fiber.App
	Repository *repository.BookRepository
}

func (s *Server) Start() {

	s.App = fiber.New()

	s.App.Get("/books/", s.GetAll)
	s.App.Get("/books/:id", s.GetBook)
	s.App.Post("/books", s.CreateBook)
	s.App.Put("/books/:id", s.UpdateBook)
	s.App.Delete("/books/:id", s.DeleteBook)

	s.App.Listen(":3000")
}
