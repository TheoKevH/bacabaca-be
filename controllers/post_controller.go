package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheoKevH/bacabaca-be/database"
	db "github.com/TheoKevH/bacabaca-be/db/generated"
	"github.com/TheoKevH/bacabaca-be/middleware"
	"github.com/TheoKevH/bacabaca-be/models"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var input models.CreatePostInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	email := r.Context().Value(middleware.UserEmailKey).(string)
	queries := db.New(database.DB)

	// Look up user ID
	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Create post
	params := db.CreatePostParams{
		Title:    input.Title,
		Slug:     input.Slug,
		Content:  input.Content,
		AuthorID: user.ID,
	}

	err = queries.CreatePost(context.Background(), params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating post: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Post created successfully")
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	queries := db.New(database.DB)

	posts, err := queries.ListPosts(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func GetPostBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	queries := db.New(database.DB)
	post, err := queries.GetPostBySlug(context.Background(), slug)
	if err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	var input models.UpdatePostInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	email := r.Context().Value(middleware.UserEmailKey).(string)
	queries := db.New(database.DB)

	// Lookup user ID from email
	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	postUUID := pgtype.UUID{}
	err = postUUID.Scan(postID)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	authorUUID := user.ID

	// Update the post (if owned by user)
	err = queries.UpdatePost(context.Background(), db.UpdatePostParams{
		Title:    input.Title,
		Content:  input.Content,
		ID:       postUUID,
		AuthorID: authorUUID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating post: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Post updated successfully")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	email := r.Context().Value(middleware.UserEmailKey).(string)
	queries := db.New(database.DB)

	user, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	postUUID := pgtype.UUID{}
	err = postUUID.Scan(postID)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	// Use user.ID directly (already pgtype.UUID)
	err = queries.DeletePost(context.Background(), db.DeletePostParams{
		ID:       postUUID,
		AuthorID: user.ID,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting post: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
