package router

import (
	"desafio_api/handler"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func NewRouter(e *echo.Echo, productHandler *handler.ProductHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products", productHandler.GetProducts)
	e.GET("/products/:id", productHandler.GetProductByID)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)
}
