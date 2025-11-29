// Package dto contiene las estructuras de datos (Data Transfer Objects)
// utilizadas para las solicitudes y respuestas de la API.
package dto

// RegisterRequest es el DTO para el endpoint de registro.
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// LoginRequest es el DTO para el endpoint de inicio de sesión.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// AuthResponse es el DTO que enviamos al cliente tras un Login o Register exitoso.
type AuthResponse struct {
	Token string `json:"token"`
	// Opcional: información básica del usuario
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}
