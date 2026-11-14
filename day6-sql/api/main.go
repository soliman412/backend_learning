package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ---------- Types ----------

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// ---------- Database ----------

var db *pgxpool.Pool

// ---------- Helpers ----------

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func parseID(r *http.Request) (int, bool) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		return 0, false
	}
	return id, true
}

func serverError(w http.ResponseWriter, err error) {
	log.Println("database error:", err)
	writeJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "internal server error"})
}

// ---------- Handlers ----------

func getEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(r.Context(), "SELECT id, name, salary FROM employees ORDER BY id")
	if err != nil {
		serverError(w, err)
		return
	}
	defer rows.Close()

	employees := []Employee{}

	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Salary)
		if err != nil {
			serverError(w, err)
			return
		}
		employees = append(employees, e)
	}

	err = rows.Err()
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

	var e Employee
	err := db.QueryRow(r.Context(),
		"SELECT id, name, salary FROM employees WHERE id = $1",
		id,
	).Scan(&e.ID, &e.Name, &e.Salary)

	if errors.Is(err, pgx.ErrNoRows) {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
		return
	}
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, e)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var input Employee

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}

	if input.Name == "" || input.Salary <= 0 {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "name is required and salary must be positive"})
		return
	}

	var e Employee
	err = db.QueryRow(r.Context(),
		"INSERT INTO employees (name, salary) VALUES ($1, $2) RETURNING id, name, salary",
		input.Name, input.Salary,
	).Scan(&e.ID, &e.Name, &e.Salary)
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, e)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

	var input Employee

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON: " + err.Error()})
		return
	}

	if input.Name == "" || input.Salary <= 0 {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "name is required and salary must be positive"})
		return
	}

	var e Employee
	err = db.QueryRow(r.Context(),
		"UPDATE employees SET name = $1, salary = $2 WHERE id = $3 RETURNING id, name, salary",
		input.Name, input.Salary, id,
	).Scan(&e.ID, &e.Name, &e.Salary)

	if errors.Is(err, pgx.ErrNoRows) {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
		return
	}
	if err != nil {
		serverError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, e)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

	result, err := db.Exec(r.Context(), "DELETE FROM employees WHERE id = $1", id)
	if err != nil {
		serverError(w, err)
		return
	}

	if result.RowsAffected() == 0 {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ---------- Main ----------

func main() {
	ctx := context.Background()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatal("cannot create pool: ", err)
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("cannot reach database: ", err)
	}

	db = pool
	fmt.Println("Connected to PostgreSQL!")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", getEmployees)
	mux.HandleFunc("GET /employees/{id}", getEmployee)
	mux.HandleFunc("POST /employees", createEmployee)
	mux.HandleFunc("PUT /employees/{id}", updateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", deleteEmployee)

	fmt.Println("Server running on http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
