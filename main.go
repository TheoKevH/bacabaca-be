package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()

	// Test route
    r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Bacabaca API is healthy!")
    }).Methods("GET")

    fmt.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}