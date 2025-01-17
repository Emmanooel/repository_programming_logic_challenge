package web

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(routes *fiber.App, server *Server) *fiber.App {
	routes.Post("/averageCalculator", func(c *fiber.Ctx) error {
		return server.AverageCalculator.AverageCalculatorHandler(c)
	})

	return routes
}
