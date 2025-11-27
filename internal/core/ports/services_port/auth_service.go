package servicesport

import (
	"context"

	"github.com/doble97/scheduleapi/internal/core/domain"
)

// AuthService es el puerto de entrada principal para las operaciones de autenticación.
// Define la funcionalidad de negocio que los Adaptadores (Handlers) usarán.
type AuthService interface {
	Register(ctx context.Context, req domain.User) (*domain.User, error)
	Login(ctx context.Context, req domain.User) (*domain.User, error)
}
