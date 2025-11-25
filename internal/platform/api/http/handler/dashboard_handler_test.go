package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	mock_service "github.com/doble97/scheduleapi/internal/platform/api/http/mocks"
	"github.com/doble97/scheduleapi/pkg/error_app"
	"go.uber.org/mock/gomock"
)

func TestDashboardHandler_CreateDashboard_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ARRANGE
	mockService := mock_service.NewMockDashboardService(ctrl)
	h := NewDashboardHandler(mockService)

	// JSON Inválido
	invalidJSON := []byte(`{"name": "Test", "email": 123`) // JSON malformado

	// El servicio NO debe ser llamado
	mockService.EXPECT().CreateDashboard(gomock.Any()).Times(0) // Esta aserción de mockgen sigue siendo válida

	// Simulación de Petición
	req := httptest.NewRequest(http.MethodPost, "/dashboards", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// ACT
	h.CreateDashboardHandler(rr, req)

	// ASSERT (Usando net/http/httptest y testing estándar)

	// 1. Verificar el Código de Estado (Reemplaza assert.Equal)
	expectedCode := http.StatusBadRequest
	if rr.Code != expectedCode {
		t.Errorf("Expected status code %d, got %d", expectedCode, rr.Code)
	}

	// 2. Verificar el Contenido del Cuerpo (Reemplaza assert.Contains)
	expectedErrorMessage := "The request body or parameters are malformed"
	responseBody := rr.Body.String()

	if !strings.Contains(responseBody, expectedErrorMessage) {
		t.Errorf("Expected body to contain '%s', but got '%s'", expectedErrorMessage, responseBody)
	}
}

func TestDashboardHandler_CreateDashboard_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ARRANGE
	mockService := mock_service.NewMockDashboardService(ctrl)
	h := NewDashboardHandler(mockService)

	// 1. DTO de Petición
	requestBody := dto.DashboardRequest{Title: "Panel de Prueba", Description: "Sirve para algo"}
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// 2. Comportamiento esperado del servicio (Core)
	expectedDomain := domain.Dashboard{ID: 1, Title: "Panel de Prueba", Description: "Sirve para algo"}
	mockService.EXPECT().
		CreateDashboard(gomock.Any()).
		Return(expectedDomain, nil).
		Times(1)

	// 3. Simulación de la Petición HTTP
	req := httptest.NewRequest(http.MethodPost, "/dashboard", bytes.NewBuffer(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	// 4. Grabadora de Respuesta
	rr := httptest.NewRecorder()

	// ACT
	h.CreateDashboardHandler(rr, req)

	// ASSERT (Aserciones Estándar)
	expectedCode := http.StatusCreated

	// Verificar Código de Estado
	if rr.Code != expectedCode {
		t.Errorf("FAIL: Expected status code %d, got %d", expectedCode, rr.Code)
	}

	// Verificar Content-Type
	expectedContentType := "application/json"
	if rr.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("FAIL: Expected Content-Type '%s', got '%s'", expectedContentType, rr.Header().Get("Content-Type"))
	}

	// Verificar el Cuerpo de la Respuesta JSON
	var responseDTO dto.DashboardResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &responseDTO); err != nil {
		t.Fatalf("FAIL: Could not unmarshal response body: %v", err)
	}

	if responseDTO.ID != 1 {
		t.Errorf("FAIL: Expected ID '1', got '%v'", responseDTO.ID)
	}
	if responseDTO.Title != "Panel de Prueba" {
		t.Errorf("FAIL: Expected Name 'Panel de Prueba', got '%s'", responseDTO.Title)
	}
}

func TestDashboardHandler_CreateDashboard_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ARRANGE
	mockService := mock_service.NewMockDashboardService(ctrl)
	h := NewDashboardHandler(mockService)

	// Petición válida, pero con datos que fallarán la validación de negocio (ej. Name vacío)
	requestBody := dto.DashboardRequest{Title: "  "}
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Comportamiento del Mock: Devuelve el error de Dominio Genérico
	mockService.EXPECT().
		CreateDashboard(gomock.Any()).
		Return(domain.Dashboard{}, domain.ErrInvalidData). // 👈 El Handler debe manejar este error
		Times(1)

	// Simulación de Petición
	req := httptest.NewRequest(http.MethodPost, "/dashboards", bytes.NewBuffer(requestBodyJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// ACT
	h.CreateDashboardHandler(rr, req)

	// ASSERT (Aserciones Estándar)
	expectedCode := http.StatusUnprocessableEntity

	// Verificar Código de Estado
	if rr.Code != expectedCode {
		t.Errorf("FAIL: Expected status code %d, got %d", expectedCode, rr.Code)
	}

	// Verificar el Cuerpo de la Respuesta JSON Estructurada
	// Nota: El handler debe devolver un JSON estructurado con el código de error
	var errResponse error_app.ErrorResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &errResponse); err != nil {
		t.Fatalf("FAIL: Could not unmarshal error response: %v", err)
	}

	// Verificar el Código de Error Interno
	expectedErrorCode := error_app.InvalidInputError.ErrorCode // e.g., "INVALID_INPUT"
	if errResponse.Code != expectedErrorCode {
		t.Errorf("FAIL: Expected error code '%s', got '%s'", expectedErrorCode, errResponse.Code)
	}

	// Verificar el Status HTTP dentro del JSON
	if errResponse.Status != expectedCode {
		t.Errorf("FAIL: Expected JSON status %d, got %d", expectedCode, errResponse.Status)
	}
}
