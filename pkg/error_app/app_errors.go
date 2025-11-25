package error_app

import "net/http"

// ErrorResponse es la estructura que se serializará a JSON
// y se enviará al cliente.
type ErrorResponse struct {
	Status int    `json:"status"` // Código de estado HTTP (400, 422, 500)
	Code   string `json:"code"`   // Código de error interno (E.g., "INVALID_INPUT", "NOT_FOUND")
	Detail string `json:"detail"` // Descripción detallada del error
}

// ErrorMapper es una estructura que contiene el error del dominio
// y su correspondiente respuesta HTTP.
type ErrorMapper struct {
	HTTPStatus int
	ErrorCode  string
	Detail     string
}

// Error implements error.
func (e ErrorMapper) Error() string {
	panic("unimplemented")
}

// Catálogo de Errores de Dominio (Centralización de Códigos de Negocio)
var (
	// Errores de Negocio/Validación
	InvalidInputError  = ErrorMapper{http.StatusUnprocessableEntity, "INVALID_INPUT", "The provided data is invalid or incomplete."}
	AlreadyExistsError = ErrorMapper{http.StatusConflict, "ALREADY_EXISTS", "The resource already exists."}

	// Errores de Adaptador/HTTP (Parsing, por ejemplo)
	BadRequestError = ErrorMapper{http.StatusBadRequest, "BAD_REQUEST", "The request body or parameters are malformed."}

	// Error Genérico de Servidor
	InternalServerError = ErrorMapper{http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected internal error occurred."}
)
