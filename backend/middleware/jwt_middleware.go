package middleware

import (
	"net/http"
	"strings"
	"warehouse/logic"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

// JWTMiddleware is a middleware that validates JWT token in request header.
func JWTMiddleware(secret string) mux.MiddlewareFunc {
	// Middleware function that processes each incoming request
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header.
			tokenString := r.Header.Get("Authorization")
			// Remove the "Bearer " prefix from the token string.
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			claims := &logic.Claims{}
			// Parse token and verify claims.
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				// Return secret key to validate token.
				return []byte(secret), nil
			})

			// validation token not valid, return an Unauthorized error.
			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
