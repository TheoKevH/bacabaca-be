package routes

import (
	"github.com/TheoKevH/bacabaca-be/controllers"
	"github.com/TheoKevH/bacabaca-be/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/register", controllers.RegisterUser).Methods("POST")
	r.HandleFunc("/api/auth/login", controllers.LoginUser).Methods("POST")
	r.HandleFunc("/api/posts", controllers.GetAllPosts).Methods("GET")
	r.HandleFunc("/api/posts/{slug}", controllers.GetPostBySlug).Methods("GET")

	auth := r.PathPrefix(("/api")).Subrouter()
	auth.Use(middleware.JWTMiddleware)
	auth.HandleFunc("/posts", controllers.CreatePost).Methods("POST")
	auth.HandleFunc("/posts/{id}", controllers.UpdatePost).Methods("PUT")
	auth.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")
}
