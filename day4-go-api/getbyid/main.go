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

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	idText := r.PathValue("id")

	id, err := strconv.Atoi(idText)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id must be a number"})
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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", getEmployees)
	mux.HandleFunc("GET /employees/{id}", getEmployee)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}