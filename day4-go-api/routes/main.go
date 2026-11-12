package main 

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request:", r.Method, r.URL.Path)
	fmt.Fprintln(w, "Welcome to the Employees API")
}

func about(w http.ResponseWriter, r*http.Request) {
	fmt.Fprintln(w, "This API is built by SOLIMAN")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /about", about)
	mux.HandleFunc("GET /health", health)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error:", err)
	}
}