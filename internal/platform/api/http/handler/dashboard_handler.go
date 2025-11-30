package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	"github.com/doble97/scheduleapi/internal/platform/api/http/util"
	"github.com/gorilla/mux"
)

type DashboardHandler struct {
	dashboardService ports.DashboardService
}

func NewDashboardHandler(dashboardService ports.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) CreateDashboardHandler(w http.ResponseWriter, r *http.Request) error {
	// NOT IMPLEMENTED
	dashReq, err := util.DecodeBodyGen[dto.DashboardRequest](r)
	if err != nil {
		return err
	}
	w.Header().Set("Content-type", "application/json")
	newDashboard := domain.Dashboard{
		Title:       dashReq.Title,
		Description: dashReq.Description,
	}
	response, err := h.dashboardService.CreateDashboard(newDashboard)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(util.DomainToDTOResponse(response)); err != nil {
		return err
	}
	return nil
}

func (h *DashboardHandler) GetManyDashboardsByIDUserHandler(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	userID := vars["userID"]
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	response, err := h.dashboardService.GetManyDashboardsByIDUser(id)
	if err != nil {
		return err
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		return err
	}
	return nil
}
