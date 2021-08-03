package main

import (
	"database/sql"

	db2 "github.com/LeonardoBatistaCarias/fc2-hexagonal-architecture/adapters/db"
	"github.com/LeonardoBatistaCarias/fc2-hexagonal-architecture/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)
}
