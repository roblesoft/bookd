package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roblesoft/bookd/pkg/models"
)

func GetBook(c *fiber.Ctx) error {

	json := &models.Book{
		Name:   "Principlesss",
		Author: "Ray Dalio",
		Year:   2016,
	}

	return c.JSON(json)
}
