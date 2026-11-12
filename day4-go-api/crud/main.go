package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"sync"
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

var mu sync.Mutex

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

func findIndex(id int) int {
	for i, e := range employees {
		if e.ID == id {
			return i
		}
	}
	return -1
}

// ---------- Handlers ----------

func getEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	writeJSON(w, http.StatusOK, employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	i := findIndex(id)
	if i == -1 {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
		return
	}

	writeJSON(w, http.StatusOK, employees[i])
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

	mu.Lock()
	defer mu.Unlock()

	input.ID = nextID
	nextID++
	employees = append(employees, input)

	writeJSON(w, http.StatusCreated, input)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

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

	mu.Lock()
	defer mu.Unlock()

	i := findIndex(id)
	if i == -1 {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
		return
	}

	employees[i].Name = input.Name
	employees[i].Salary = input.Salary

	writeJSON(w, http.StatusOK, employees[i])
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(r)
	if !ok {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a positive number"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	i := findIndex(id)
	if i == -1 {
		writeJSON(w, http.StatusNotFound, ErrorResponse{Error: "employee not found"})
		return
	}

	employees = slices.Delete(employees, i, i+1)

	w.WriteHeader(http.StatusNoContent)
}

// ---------- Main ----------

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", getEmployees)
	mux.HandleFunc("GET /employees/{id}", getEmployee)
	mux.HandleFunc("POST /employees", createEmployee)
	mux.HandleFunc("PUT /employees/{id}", updateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", deleteEmployee)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}