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
	var err error
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	num, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal("Error connect to the database:", err)
	}
	err = num.Ping()
	if err != nil {
		log.Fatal("failed to ping ", err)
	}

	return num
}

type product struct {
	product_name  string
	unit          string
	category_name string
	price         float32
	description   string
}

func main() {

	num := connect()
	defer num.Close()

	sql := `SELECT p.product_name, p.unit, c.category_name, p.price, c.description
	FROM pro p
	JOIN categories c ON p.category_id = c.category_id
	WHERE c.category_name = 'Beverages';`

	rows, err := num.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var products []product
	for rows.Next() {
		var n product
		err = rows.Scan(&n.product_name, &n.unit, &n.category_name, &n.price, &n.description)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, n)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
