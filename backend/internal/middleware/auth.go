package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/amilcar-vasquez/auth-service/backend/internal/utils"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

const UserIDKey contextKey = "userID"

// AuthMiddleware validates JWT tokens and protects routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		// Extract token from "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid authorization header format")
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserIDFromContext extracts the user ID from the request context
func GetUserIDFromContext(r *http.Request) (uint, bool) {
	userID, ok := r.Context().Value(UserIDKey).(uint)
	return userID, ok
}
