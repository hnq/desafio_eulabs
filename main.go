package main

import (
	"desafio_api/config"
	"desafio_api/handler"
	"desafio_api/repository"
	"desafio_api/router"
	"desafio_api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inicializar banco de dados
	config.InitDB()
	defer config.DB.Close()

	// Criar instâncias de repositório, serviço e handler
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	// Inicializar Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configurar rotas
	router.NewRouter(e, productHandler)

	// Iniciar servidor
	e.Logger.Fatal(e.Start(":8080"))
}
