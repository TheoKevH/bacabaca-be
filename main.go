package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheoKevH/bacabaca-be/database"
	"github.com/TheoKevH/bacabaca-be/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Bacabaca API is healthy!")
	}).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
