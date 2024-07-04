package service

import (
	"desafio_api/model"
	"desafio_api/repository"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProducts() ([]model.Product, error)
	GetProductByID(id int) (*model.Product, error)
	UpdateProduct(product *model.Product) error
	DeleteProduct(id int) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetProducts() ([]model.Product, error) {
	return s.repo.GetAll()
}

func (s *productService) GetProductByID(id int) (*model.Product, error) {
	return s.repo.GetByID(id)
}

func (s *productService) UpdateProduct(product *model.Product) error {
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id int) error {
	return s.repo.Delete(id)
}
