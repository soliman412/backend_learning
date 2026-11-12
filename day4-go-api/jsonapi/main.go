package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Salary float64 `json:"salary"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request:", r.Method, r.URL.Path)

	employees :=[]Employee {
		{ID: 1, Name: "Ahmad", Salary: 3000},
		{ID: 2, Name: "Sara", Salary: 1200},
		{ID: 3, Name: "Soliman", Salary: 6000},				
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", getEmployees)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}	
}