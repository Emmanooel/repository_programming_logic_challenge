package handlers

import (
	dto "desafio_primeira_semana/internal/DTO"
	exceptions "desafio_primeira_semana/internal/Exceptions"

	"github.com/gofiber/fiber/v2"
)

type AverageCalculatorHandlerInterface interface {
	AverageCalculatorHandler(c *fiber.Ctx) error
}

type AverageCalculatorHandlerStruct struct {
}

func NewAverageCalculatorHandler() AverageCalculatorHandlerInterface {
	return AverageCalculatorHandlerStruct{}
}

func (a AverageCalculatorHandlerStruct) AverageCalculatorHandler(c *fiber.Ctx) error {
	var averageCalculatorDTO dto.AverageCalculatorDTO

	body := c.BodyParser(&averageCalculatorDTO)

	if body != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": exceptions.AverageCalculatorBodyParseError})
	}

	average := (averageCalculatorDTO.Number1 + averageCalculatorDTO.Number2 + averageCalculatorDTO.Number3) / 3

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"average": average})
}
