package middleware

import (
	"fmt"
	"go-backend/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Middleware HIT")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "Missing token", 401)
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid token format", 401)
			return
		}

		tokenStr := parts[1]

		_, err := utils.VerifyToken(tokenStr)
		if err != nil {
			fmt.Println("JWT error:", err)
			http.Error(w, "Invalid token", 401)
			return
		}

		fmt.Println("Token valid")

		next(w, r)
	}
}