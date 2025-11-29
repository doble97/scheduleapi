package services

import (
	"context"
	"errors"

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
	// TODO: Mejorar el manejo de errores para enviar  codigos HTTP más adecuados
	existingUser, err := a.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, domain.ErrorInternalServer
	}
	if existingUser == nil {
		return nil, domain.ErrorInternalServer
	}
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.Password))
	if err != nil {
		return nil, domain.ErrorInternalServer
	}

	return existingUser, nil
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
	idUser, err := a.userRepo.SaveUser(ctx, user)
	if err != nil {
		return nil, domain.ErrorInternalServer
	}
	user.ID = idUser
	return &user, nil
}

func NewAuthService(u ports.UserRepository) servicesport.AuthService {
	return &authService{
		userRepo: u,
	}
}
