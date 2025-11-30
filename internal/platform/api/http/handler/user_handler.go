// Package handler is for routes
package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/doble97/scheduleapi/internal/core/domain"
	servicesport "github.com/doble97/scheduleapi/internal/core/ports/services_port"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	"github.com/doble97/scheduleapi/internal/platform/api/http/util"
)

type UserHandler struct {
	service servicesport.AuthService
}

func NewUserHandler(servi servicesport.AuthService) *UserHandler {
	return &UserHandler{
		service: servi,
	}
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) error {
	loginReq, err := util.DecodeBodyGen[dto.LoginRequest](r)
	if err != nil {
		return err
	}
	user := domain.User{Email: loginReq.Email, Password: loginReq.Password}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := h.service.Login(ctx, user)
	if err != nil {
		return err
	}
	token, err := util.GenerateJWT(*result)
	if err != nil {
		return err
	}
	authResponse := util.UserDomainToAuthResponse(*result, token)

	json.NewEncoder(w).Encode(authResponse)
	return nil
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) error {
	registerReq, err := util.DecodeBodyGen[dto.RegisterRequest](r)
	if err != nil {
		return err
	}
	user := domain.User{Name: registerReq.Name, LastName: registerReq.LastName, Email: registerReq.Email, Password: registerReq.Password}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := h.service.Register(ctx, user)
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(result)
	return nil
}
