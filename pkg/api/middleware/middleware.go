package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware defines middleware for jwt authentication
type JWTMiddleware struct {
	secret string
}

// NewJWTMiddleware returns a new JWTMiddleware
func NewJWTMiddleware(secret string) *JWTMiddleware {
	return &JWTMiddleware{secret}
}

// AuthHandler returns an http handler that performs jwt authentication
func (j *JWTMiddleware) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		// Check and strip "Bearer " prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization Header Format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Ensure the token is signed using the expected method
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(j.secret), nil
		})
		if err != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r) // Pass request to the next handler
	})
}
