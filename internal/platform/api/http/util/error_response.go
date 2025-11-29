package util

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/doble97/scheduleapi/pkg/error_app"
)

func RespondWithError(w http.ResponseWriter, err error) {
	mappedError := error_app.MapErrorToHTTP(err)
	if mappedError.HTTPStatus >= 500 {
		log.Printf("ERROR 500: %v", err)
	}
	response := error_app.ErrorResponse{
		Status: mappedError.HTTPStatus,
		Code:   mappedError.ErrorCode,
		Detail: mappedError.Detail,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(mappedError.HTTPStatus)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error al codificar la respuesta de error: %v", err)
	}
}
