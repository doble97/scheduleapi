// Package middleware is package for middlewares
package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
)

type JWTVerifier func(tokenString string) (*dto.CustomClaims, error)

func AuthMiddlewareFactory(verifyFunc JWTVerifier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				errorResponse := struct {
					Error   string `json:"error"`
					Message string `json:"message"`
				}{
					Error:   "Unauthorized",
					Message: "Token de autorización requerido",
				}
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
					http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
				}
				return
			}
			_, err := verifyFunc(token)
			if err != nil {
				http.Error(w, "Token no válido", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
