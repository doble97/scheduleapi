package domain

import "errors"

// ⚠️ ERRORES GENÉRICOS DE DOMINIO ⚠️
// Usados por cualquier Servicio o Entidad del Core.

var (
	// Se usa cuando un recurso específico no existe.
	ErrNotFound = errors.New("resource not found")

	// Se usa cuando una entidad que se intenta crear ya existe (e.g., email/nombre duplicado).
	ErrAlreadyExists = errors.New("resource already exists")

	// Se usa cuando los datos de entrada violan las reglas de negocio/validación.
	ErrInvalidData = errors.New("invalid data provided")

	// Se usa para indicar que el usuario o contexto no tiene permiso para la acción.
	ErrForbidden = errors.New("operation forbidden")

	// Se usa para problemas de concurrencia o de bloqueo.
	ErrConflict = errors.New("conflict occurred during update")
)
