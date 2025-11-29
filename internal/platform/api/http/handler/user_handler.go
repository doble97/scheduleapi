// Package handler is for routes
package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/doble97/scheduleapi/internal/core/domain"
	servicesport "github.com/doble97/scheduleapi/internal/core/ports/services_port"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	"github.com/doble97/scheduleapi/internal/platform/api/http/util"
	"github.com/go-playground/validator/v10"
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
	var loginReq dto.LoginRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &loginReq); err != nil {
		return err
	}
	if err := validator.New().Struct(loginReq); err != nil {
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
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(authResponse)
	return nil
}

func (h *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) error {
	var registerReq dto.RegisterRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &registerReq); err != nil {
		return err
	}
	if err := validator.New().Struct(registerReq); err != nil {
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
