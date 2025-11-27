package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
	servicesport "github.com/doble97/scheduleapi/internal/core/ports/services_port"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo ports.UserRepository
}

// Login implements servicesport.AuthService.
func (a *authService) Login(ctx context.Context, req domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// Register implements servicesport.AuthService.
func (a *authService) Register(ctx context.Context, req domain.User) (*domain.User, error) {
	existingUser, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return nil, domain.ErrorInternalServer
	}
	if existingUser != nil {
		return nil, domain.ErrorInternalServer
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, domain.ErrorInternalServer
	}
	user := domain.User{
		Name:     req.Name,
		LastName: req.LastName,
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	token := "JWT_GENERADO" + req.Email
	fmt.Println("TOKEN---", token)
	return &user, nil

}

func NewAuthService(u ports.UserRepository) servicesport.AuthService {
	return &authService{
		userRepo: u,
	}
}
