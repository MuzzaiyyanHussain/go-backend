package handlers

import (
	"encoding/json"
	"go-backend/db"
	"go-backend/models"
	"net/http"
)

// GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var id int
	err = db.DB.QueryRow(
		"INSERT INTO users(name) VALUES($1) RETURNING id",
		user.Name,
	).Scan(&id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	user.ID = id
	json.NewEncoder(w).Encode(user)
}