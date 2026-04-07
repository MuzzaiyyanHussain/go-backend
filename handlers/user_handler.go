package handlers

import (
	"encoding/json"
	"go-backend/db"
	"go-backend/models"
	"go-backend/utils"
	"net/http"
)

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

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	var dbUser models.User

	err := db.DB.QueryRow(
		"SELECT id, name FROM users WHERE name=$1",
		user.Name,
	).Scan(&dbUser.ID, &dbUser.Name)

	if err != nil {
		http.Error(w, "User not found", 401)
		return
	}

	token, _ := utils.GenerateToken(dbUser.ID)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}