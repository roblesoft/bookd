package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	usecase "github.com/roblesoft/bookd/internal/usecase"
)

type HTTP struct {
	Port    string
	App     *fiber.App
	Service *usecase.Service
}

func (s *HTTP) Start() {
	s.App = fiber.New()
	s.App.Use(logger.New())
	s.App.Get("/books/", s.GetAll)
	s.App.Get("/books/:id", s.GetBook)
	s.App.Post("/books", s.CreateBook)
	s.App.Put("/books/:id", s.UpdateBook)
	s.App.Delete("/books/:id", s.DeleteBook)
	s.App.Listen(":3000")
}
