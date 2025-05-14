package routes

import (
	"github.com/TheoKevH/bacabaca-be/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/auth/login", controllers.LoginUser).Methods("POST")
}
