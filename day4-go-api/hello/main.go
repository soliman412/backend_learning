package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go server!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", hello)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
