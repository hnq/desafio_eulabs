package service

import (
	"testing"

	"desafio_api/model"

	"github.com/stretchr/testify/assert"
)

type mockProductRepository struct{}

func (m *mockProductRepository) Create(product *model.Product) error {
	return nil
}

func (m *mockProductRepository) GetAll() ([]model.Product, error) {
	products := []model.Product{
		{ID: 1, Name: "Product 1", Description: "Description 1", Price: 19.99},
		{ID: 2, Name: "Product 2", Description: "Description 2", Price: 29.99},
	}
	return products, nil
}

func (m *mockProductRepository) GetByID(id int) (*model.Product, error) {
	product := &model.Product{ID: id, Name: "Product 1", Description: "Description 1", Price: 19.99}
	return product, nil
}

func (m *mockProductRepository) Update(product *model.Product) error {
	return nil
}

func (m *mockProductRepository) Delete(id int) error {
	return nil
}

func TestProductService(t *testing.T) {
	repo := &mockProductRepository{}
	service := NewProductService(repo)

	t.Run("GetAllProducts", func(t *testing.T) {
		products, err := service.GetProducts()
		assert.NoError(t, err)
		assert.NotEmpty(t, products)
	})

	t.Run("GetProductByID", func(t *testing.T) {
		product, err := service.GetProductByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, product)
	})
}
