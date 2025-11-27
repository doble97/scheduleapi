package handler

import (
	"net/http"

	"github.com/doble97/scheduleapi/internal/platform/api/http/util"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func ErrorHandlerMiddleware(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			util.RespondWithError(w, err)
			return
		}
	}
}
