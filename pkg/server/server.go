package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesoft/bookd/pkg/routes"
)

type Server struct {
	Port string
	App  *fiber.App
}

func (s *Server) Start() {

	s.App = fiber.New()

	s.App.Get("/books/:id", routes.GetBook)

	s.App.Listen(":3000")
}
