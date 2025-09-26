package middleware

import (
	"context"
	"nba-api/internal/response"
	"net/http"
	"strings"
)

// Simple Bearer token list
var validTokens = map[string]string{
	"secret123": "LeBron",
}

func BearerAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseWithError(w, http.StatusUnauthorized, "missing auth header")
			return
		}

		// Checking format: "Bearer <token"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.ResponseWithError(w, http.StatusUnauthorized, "invalid auth header format")
			return
		}

		token := parts[1]

		user, ok := validTokens[token]
		if !ok {
			response.ResponseWithError(w, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
