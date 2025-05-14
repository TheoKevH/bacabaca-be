package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheoKevH/bacabaca-be/database"
	db "github.com/TheoKevH/bacabaca-be/db/generated"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Prepare and call sqlc-generated CreateUser
	queries := db.New(database.DB)

	params := db.CreateUserParams{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}

	err = queries.CreateUser(context.Background(), params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User registered successfully")
}
