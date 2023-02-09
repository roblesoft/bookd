package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesoft/bookd/pkg/models"
)

func (s *Server) GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	item, err := s.Repository.Get(id)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(item)
}

func (s *Server) CreateBook(c *fiber.Ctx) error {
	var book models.Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	if err := s.Repository.Create(&book); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(book)
}
