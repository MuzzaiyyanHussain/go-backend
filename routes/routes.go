package routes

import (
	"go-backend/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handlers.GetUsers(w, r)
		case "POST":
			handlers.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", 405)
		}
	})
}