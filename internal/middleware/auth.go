package middleware

import (
	"context"
	"fmt"
	"nba-api/internal/response"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Simple Bearer token list
var validTokens = map[string]string{
	"secret123": "LeBron",
}

func GenerateJWT(subject string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", fmt.Errorf("JWT_SECRET not set")
	}

	claims := jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Minute * 15).Unix(), // 30 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
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

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseWithError(w, http.StatusUnauthorized, "missing auth header")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.ResponseWithError(w, http.StatusUnauthorized, "invalid auth header format")
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				return nil, fmt.Errorf("JWT_SECRET env variable not set")
			}
			return []byte(secret), nil
		})
		if err != nil {
			response.ResponseWithError(w, http.StatusUnauthorized, "failed to parse token: "+err.Error())
			return
		}

		// Extract claims safely
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			response.ResponseWithError(w, http.StatusUnauthorized, "invalid token claims")
			return
		}

		// Check sub
		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			response.ResponseWithError(w, http.StatusUnauthorized, "token missing sub claim")
			return
		}

		// Pass identity in context
		ctx := context.WithValue(r.Context(), "user", sub)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
