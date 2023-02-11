package controllers

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/roblesoft/bookd/internal/entity"
)

func (s *HTTP) GetAll(c *fiber.Ctx) error {
	limit := c.Query("limit")
	page := c.Query("page")
	items, err := s.Repository.GetAll(limit, page)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(items)
}

func (s *HTTP) GetBook(c *fiber.Ctx) error {
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

func (s *HTTP) CreateBook(c *fiber.Ctx) error {
	var book entity.Book
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

func (s *HTTP) UpdateBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	book.ID = id
	if err := s.Repository.Update(&book); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(book)
}

func (s *HTTP) DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := s.Repository.Delete(id); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.SendStatus(204)
}
