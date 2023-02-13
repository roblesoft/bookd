package controllers

import (
	"github.com/gofiber/fiber/v2"
	entity "github.com/roblesoft/bookd/internal/entity"
)

func (h *HTTP) GetAll(c *fiber.Ctx) error {
	limit := c.Query("limit")
	page := c.Query("page")
	items, err := h.Service.GetAllBooks(limit, page)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(items)
}

func (h *HTTP) GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	item, err := h.Service.GetBook(id)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(item)
}

func (h *HTTP) CreateBook(c *fiber.Ctx) error {
	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	if err := h.Service.CreateBook(&book); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(book)
}

func (h *HTTP) UpdateBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	var book entity.Book
	if err := c.BodyParser(&book); err != nil {
		c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	book.ID = id
	if err := h.Service.UpdateBook(&book); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.JSON(book)
}

func (h *HTTP) DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.Service.DeleteBook(id); err != nil {
		c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
		return err
	}
	return c.SendStatus(204)
}
