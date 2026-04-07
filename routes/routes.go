package routes

import (
	"go-backend/handlers"
	"go-backend/middleware"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/login", handlers.Login)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case "POST":
			handlers.CreateUser(w, r)

		case "GET":
			middleware.AuthMiddleware(handlers.GetUsers)(w, r)

		default:
			http.Error(w, "Method not allowed", 405)
		}
	})
}