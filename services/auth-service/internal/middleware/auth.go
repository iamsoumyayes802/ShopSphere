package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey contextKey = "role"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHandler := r.Header.Get("Authorization")

		if authHandler 	== "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return 
		}

		parts := strings.Split(authHandler, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header", http.StatusUnauthorized)
			return 
		}

		tokenString := parts[1]

		token, err := jwt.Parse(
			tokenString,
			func (token *jwt.Token) (interface{}, error)  {
				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)

		if !ok {
			http.Error(w, "invalid user id", http.StatusUnauthorized)
			return
		}

		role, _ := claims["role"].(string)

		ctx := context.WithValue(
			r.Context(),
			UserIDKey,
			userID,
		)

		ctx = context.WithValue(
			ctx,
			RoleKey,
			role,
		)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}