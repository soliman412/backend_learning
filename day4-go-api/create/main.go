package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var employees = []Employee{
	{ID: 1, Name: "Ahmad", Salary: 3000},
	{ID: 2, Name: "Sara", Salary: 1200},
	{ID: 3, Name: "Soliman", Salary: 6000},
}

var nextID = 4

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

	for _, e := range employees {
		if e.ID == id {
			writeJSON(w, http.StatusOK, e)
			return
		}
	}

	writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var input Employee

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid JSON"})
		return
	}

	if input.Name == "" || input.Salary <= 0 {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "name is required and salary must be positive"})
		return
	}

	input.ID = nextID
	nextID++

	employees = append(employees, input)

	writeJSON(w, http.StatusCreated, input)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", getEmployees)
	mux.HandleFunc("GET /employees/{id}", getEmployee)
	mux.HandleFunc("POST /employees", createEmployee)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}