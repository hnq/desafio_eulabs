package repository

import (
	"database/sql"
	"desafio_api/model"
)

type ProductRepository interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id int) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) Create(product *model.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)", product.Name, product.Description, product.Price)
	return err
}

func (r *productRepository) GetAll() ([]model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, created_at FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *productRepository) GetByID(id int) (*model.Product, error) {
	var product model.Product
	err := r.db.QueryRow("SELECT id, name, description, price, created_at FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *model.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?", product.Name, product.Description, product.Price, product.ID)
	return err
}

func (r *productRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = ?", id)
	return err
}
