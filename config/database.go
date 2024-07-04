package config

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    // Carrega as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }

    // Pega as variáveis de ambiente
    username := os.Getenv("DB_USERNAME")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    // Cria a string DSN
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    if err = DB.Ping(); err != nil {
        panic(err)
    }
    fmt.Println("Connected to the database successfully!")
}
