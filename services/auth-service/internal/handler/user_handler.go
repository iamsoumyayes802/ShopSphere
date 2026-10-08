package handler

import (
	"auth-service/internal/middleware"
	"encoding/json"
	"net/http"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(
		middleware.UserIDKey,
	)

	role := r.Context().Value(
		middleware.RoleKey,
	)

	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userID,
		"role": role,
	})
}