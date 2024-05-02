package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Dilshod@2005"
	dbname   = "dilshod"
)

func connect() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	return db
}

type product struct {
	ProductName  string
	Unit         string
	CategoryName string
	Price        float32
	Desc         string
}

func main() {
	db := connect()
	defer db.Close()

	sql := `
		SELECT p.product_name, p.unit, c.category_name, p.price, c.Desc
		FROM products p
		JOIN categories c ON p.category_id = c.category_id
		WHERE c.category_name = 'Beverages';
	`

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []product
	for rows.Next() {
		var p product
		err := rows.Scan(&p.ProductName, &p.Unit, &p.CategoryName, &p.Price, &p.Desc)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}

	for _, p := range products {
		fmt.Printf("Product: %s\n", p.ProductName)
		fmt.Printf("  Category: %s\n", p.CategoryName)
		fmt.Printf("  Unit: %s\n", p.Unit)
		fmt.Printf("  Price: %f\n", p.Price)
		fmt.Printf("  Desc: %s\n", p.Desc)
		fmt.Println()
	}
}
