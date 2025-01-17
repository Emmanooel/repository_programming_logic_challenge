package web

import (
	handlers "desafio_primeira_semana/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App               *fiber.App
	AverageCalculator handlers.AverageCalculatorHandlerInterface
}

func NewServer(
	AverageCalculatorHandler handlers.AverageCalculatorHandlerInterface,
) *Server {
	//recebe os handlers para retornar nas rotas
	router := fiber.New()

	//inicia o server com as variaveis
	server := &Server{
		AverageCalculator: AverageCalculatorHandler,
	}

	//chama metodo das rotas
	server.App = Routes(router, server)

	return server

}
