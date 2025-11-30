// Package middleware is package for middlewares
package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
)

var mockHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Acceso concedido"))
})

func mockSuccess(tokenString string) (*dto.CustomClaims, error) {
	return &dto.CustomClaims{}, nil
}

func mockFailure(tokenString string) (*dto.CustomClaims, error) {
	return nil, errors.New("fallo de verificación JWT")
}

func TestAuthMiddleware(t *testing.T) {
	type jsonError struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}
	tests := []struct {
		name           string
		authHeader     string
		verfierFunc    JWTVerifier
		expectedStatus int
		expectedBody   string
		expectedJSON   jsonError
		contentType    string
	}{
		{
			name:           "Caso 1: Sin token (Fallo JSON 401)",
			authHeader:     "",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "", // El cuerpo es JSON
			expectedJSON: jsonError{
				Error:   "Unauthorized",
				Message: "Token de autorización requerido",
			},
			contentType: "application/json",
		},
		{
			name:           "Caso 2: Token inválido (Fallo texto 401)",
			authHeader:     "Bearer INVALID_TOKEN", // Simulación de token que falla VerifyJWT
			verfierFunc:    mockFailure,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "Token no válido\n",
			contentType:    "text/plain; charset=utf-8",
		},
		{
			name:           "Caso 3: Token válido (Éxito 200)",
			authHeader:     "Bearer VALID_TOKEN", // Simulación de token que pasa VerifyJWT
			verfierFunc:    mockSuccess,
			expectedStatus: http.StatusOK,
			expectedBody:   "Acceso concedido",
			contentType:    "text/plain; charset=utf-8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			rr := httptest.NewRecorder()
			authMiddleware := AuthMiddlewareFactory(tt.verfierFunc)
			handler := authMiddleware(mockHandler)
			handler.ServeHTTP(rr, req)

			// 1. Verificar el código de estado HTTP
			if rr.Code != tt.expectedStatus {
				t.Errorf("Código de estado esperado %d, obtenido %d", tt.expectedStatus, rr.Code)
			}

			// 2. Verificar el Content-Type
			if rr.Header().Get("Content-Type") != tt.contentType {
				t.Errorf("Content-Type esperado '%s', obtenido '%s'", tt.contentType, rr.Header().Get("Content-Type"))
			}

			// 3. Verificar el cuerpo de la respuesta
			if tt.expectedJSON.Error != "" {
				// Caso de respuesta JSON (Caso 1)
				var actualJSON jsonError
				err := json.Unmarshal(rr.Body.Bytes(), &actualJSON)
				if err != nil {
					t.Fatalf("Error al decodificar JSON: %v. Respuesta obtenida: %s", err, rr.Body.String())
				}
				if actualJSON.Error != tt.expectedJSON.Error || actualJSON.Message != tt.expectedJSON.Message {
					t.Errorf("JSON esperado %v, obtenido %v", tt.expectedJSON, actualJSON)
				}
			} else {
				// Caso de respuesta de texto (Casos 2 y 3)
				if rr.Body.String() != tt.expectedBody {
					t.Errorf("Cuerpo esperado '%s', obtenido '%s'", tt.expectedBody, rr.Body.String())
				}
			}
		})
	}
}
