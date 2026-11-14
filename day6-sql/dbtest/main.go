package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		log.Fatal("cannot connect: ", err)
	}
	defer conn.Close(ctx)

	fmt.Println("Connected to PostgreSQL!")

	// 1. Insert one employee and get its new ID
	var newID int
	err = conn.QueryRow(ctx,
		"INSERT INTO employees (name, salary) VALUES ($1, $2) RETURNING id",
		"Yara", 2800,
	).Scan(&newID)
	if err != nil {
		log.Fatal("insert failed: ", err)
	}
	fmt.Println("Inserted employee with ID:", newID)

	// 2. Read all employees
	rows, err := conn.Query(ctx, "SELECT id, name, salary FROM employees ORDER BY id")
	if err != nil {
		log.Fatal("query failed: ", err)
	}
	defer rows.Close()

	fmt.Println("All employees:")
	for rows.Next() {
		var id int
		var name string
		var salary float64

		err := rows.Scan(&id, &name, &salary)
		if err != nil {
			log.Fatal("scan failed: ", err)
		}
		fmt.Println(id, name, salary)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("rows error: ", err)
	}
}