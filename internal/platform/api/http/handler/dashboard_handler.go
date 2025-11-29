package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	"github.com/doble97/scheduleapi/internal/platform/api/http/util"
	"github.com/doble97/scheduleapi/pkg/error_app"
	"github.com/go-playground/validator/v10"
)

type DashboardHandler struct {
	dashboardService ports.DashboardService
}

func (h *DashboardHandler) respondWithError(w http.ResponseWriter, err error) {
	// var mappedError apperrors.ErrorMapper
	mappedError := error_app.MapErrorToHTTP(err)

	if mappedError.HTTPStatus >= 500 {
		log.Printf("ERROR 500: %v", err)
	}
	// 3. Crear la Respuesta JSON Estructurada
	response := error_app.ErrorResponse{
		Status: mappedError.HTTPStatus,
		Code:   mappedError.ErrorCode,
		Detail: mappedError.Detail,
	}

	// 4. Escribir la Respuesta HTTP/JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(mappedError.HTTPStatus)
	json.NewEncoder(w).Encode(response) // Ignoramos el error de Encode por simplicidad
}

func NewDashboardHandler(dashboardService ports.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

func (h *DashboardHandler) CreateDashboardHandler(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED
	w.Header().Set("Content-type", "application/json")
	var dashboard dto.DashboardRequest
	body, errBody := io.ReadAll(r.Body)
	if errBody != nil {
		w.WriteHeader(http.StatusInternalServerError)
		h.respondWithError(w, errBody)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &dashboard); err != nil {
		h.respondWithError(w, err)
		return
	}
	if err := validator.New().Struct(dashboard); err != nil {
		h.respondWithError(w, err)
		return
	}
	newDashboard := domain.Dashboard{
		Title:       dashboard.Title,
		Description: dashboard.Description,
	}
	response, err := h.dashboardService.CreateDashboard(newDashboard)
	if err != nil {
		log.Println("Error create:", err)
		h.respondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(util.DomainToDTOResponse(response))
}

func (h *DashboardHandler) GetManyDashboardsByIDUserHandler(w http.ResponseWriter, r *http.Request) {
	// NOT IMPLEMENTED
	w.Header().Set("Content-type", "application/json")
	// idUser, errParse := strconv.Atoi(chi.URLParam(r, "idUser"))
	// if errParse != nil {
	// 	h.respondWithError(w, errParse)
	// 	return
	// }
	// response, err := h.dashboardService.GetManyDashboardsByIDUser(idUser)
	response, err := h.dashboardService.GetManyDashboardsByIDUser(1) // temporal until we have auth
	if err != nil {
		log.Println("Error get many:", err)
		h.respondWithError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
