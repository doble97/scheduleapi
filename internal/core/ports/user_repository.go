package ports

import (
	"context"

	"github.com/doble97/scheduleapi/internal/core/domain"
)

// UserRepository es el puerto de salida para la persistencia de usuarios.
// Define las operaciones que el Service necesita para interactuar con la DB.
type UserRepository interface {
	// Guarda un nuevo usuario en la persistencia.
	SaveUser(ctx context.Context, user domain.User) (int, error)
	// Busca un usuario por email para el proceso de login.
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}
