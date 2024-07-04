package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"desafio_api/model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockProductService struct{}

func (m *mockProductService) CreateProduct(product *model.Product) error {
	return nil
}

func (m *mockProductService) GetProducts() ([]model.Product, error) {
	products := []model.Product{
		{ID: 1, Name: "Product 1", Description: "Description 1", Price: 19.99},
		{ID: 2, Name: "Product 2", Description: "Description 2", Price: 29.99},
	}
	return products, nil
}

func (m *mockProductService) GetProductByID(id int) (*model.Product, error) {
	product := &model.Product{ID: id, Name: "Product 1", Description: "Description 1", Price: 19.99}
	return product, nil
}

func (m *mockProductService) UpdateProduct(product *model.Product) error {
	return nil
}

func (m *mockProductService) DeleteProduct(id int) error {
	return nil
}

func TestProductHandler(t *testing.T) {
	service := &mockProductService{}
	handler := NewProductHandler(service)

	// Teste para buscar todos os produtos
	t.Run("GetProducts", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, handler.GetProducts(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Product 1")
		}
	})

	// Teste para buscar um produto por ID
	t.Run("GetProductByID", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/products/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, handler.GetProductByID(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "Product 1")
		}
	})
}
