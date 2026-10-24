package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey contextKey = "role"
)

func Auth(next http.Handler) http.Handler {
	keyBytes, err := os.ReadFile(os.Getenv("JWT_PUBLIC_KEY_ID"))
	if err != nil {
		panic("cannot load JWT public key: "+ err.Error())
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		panic("invalid JWT public key: "+ err.Error())
	}

	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		header := strings.Fields(r.Header.Get("Authorization"))
		if len(header) != 2 || header[0] != "Bearer" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(
			header[1],
			claims,
			func (token *jwt.Token) (any, error)  {
				if token.Method != jwt.SigningMethodRS256 {
					return nil, jwt.ErrSignatureInvalid
				}
				return publicKey, nil
			},
			jwt.WithValidMethods([]string{"RS256"}),
			jwt.WithIssuer("shopsphere-auth"),
			jwt.WithAudience("shopsphere"),
			jwt.WithExpirationRequired(),
		)

		if err != nil || !token.Valid {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		userID, ok := claims["sub"].(string)
		if !ok || userID == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		role, ok := claims["role"].(string)

		if !ok || (role != "USER" && role != "ADMIN") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}