package repository

import (
	"database/sql"
	"testing"

	"desafio_api/model"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {
	// Configurar conexão com o banco de dados
	// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/productdb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// Configurar conexão com o banco de dados SQLite em memória
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	// Inicializar o esquema do banco de dados, se necessário
	db.Exec("CREATE TABLE products (id INTEGER PRIMARY KEY, name TEXT, description TEXT, price DECIMAL)")

	// Criar instância do repositório
	repo := NewProductRepository(db)

	// Teste para criar um produto
	t.Run("CreateProduct", func(t *testing.T) {
		product := &model.Product{
			Name:        "Test Product",
			Description: "Test Description",
			Price:       99.99,
		}
		err := repo.Create(product)
		assert.NoError(t, err)
		assert.NotZero(t, product.ID)
	})

	// Teste para buscar todos os produtos
	t.Run("GetAllProducts", func(t *testing.T) {
		products, err := repo.GetAll()
		assert.NoError(t, err)
		assert.NotEmpty(t, products)
	})

	// Teste para buscar um produto por ID
	t.Run("GetProductByID", func(t *testing.T) {
		product, err := repo.GetByID(1)
		assert.NoError(t, err)
		assert.NotNil(t, product)
	})

	// Teste para atualizar um produto
	t.Run("UpdateProduct", func(t *testing.T) {
		updatedProduct := &model.Product{
			ID:          1,
			Name:        "Updated Product",
			Description: "Updated Description",
			Price:       129.99,
		}
		err := repo.Update(updatedProduct)
		assert.NoError(t, err)
	})

	// Teste para excluir um produto
	t.Run("DeleteProduct", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NoError(t, err)
	})
}
