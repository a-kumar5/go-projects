package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	connStr := "postgresql://postgres:postgres@localhost:5433/product?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createProductTable(db)
	product := Product{"Book", 199.55, true}
	pk := insertProduct(db, product)

	fmt.Printf("ID = %d\n", pk)

	p := queryProduct(db, pk)

	fmt.Printf("name = %s\n price = %v\n available = %v", p.Name, p.Price, p.Available)
}

func createProductTable(db *sql.DB) {
	/*
	   	Product Table

	   - ID
	   - Name
	   - Price
	   - Available
	   - Date Created
	*/

	query := `CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price NUMERIC(6,2) NOT NULL,
		available BOOLEAN,
		created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
		VALUES ( $1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}

func queryProduct(db *sql.DB, pk int) Product {
	query := "SELECT name, price, available FROM product WHERE id=$1"

	var p Product
	err := db.QueryRow(query, pk).Scan(&p.Name, &p.Price, &p.Available)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
