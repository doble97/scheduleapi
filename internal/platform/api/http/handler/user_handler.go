package handler

import (
	"encoding/json"
	"io"
	"net/http"

	servicesport "github.com/doble97/scheduleapi/internal/core/ports/services_port"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	s servicesport.AuthService
}

func NewUserHandler(servi servicesport.AuthService) *UserHandler {
	return &UserHandler{
		s: servi,
	}
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
	json.NewEncoder(w).Encode(registerReq)
	return nil
}
