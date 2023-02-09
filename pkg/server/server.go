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

	s.App.Get("/books/:id", s.GetBook)
	s.App.Post("/books", s.CreateBook)

	s.App.Listen(":3000")
}
