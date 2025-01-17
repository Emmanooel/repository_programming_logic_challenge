package internal

import (
	web "desafio_primeira_semana/infra/web"
	handlers "desafio_primeira_semana/internal/handlers"
)

func NewApplication() {
	averageCalculatorServer := handlers.NewAverageCalculatorHandler()
	srv := web.NewServer(averageCalculatorServer)

	srv.App.Listen(":3000")
}
