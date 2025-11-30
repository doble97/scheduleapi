package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func DecodeBodyGen[T any](r *http.Request) (T, error) {
	var result T
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return result, fmt.Errorf("error al leer el body: %w", err)
	}
	defer r.Body.Close() // TODO: check if it is necesary close here or out this function
	if err := json.Unmarshal(body, &result); err != nil {
		return result, fmt.Errorf("error al codificar JSON: %w", err)
	}
	if err := validator.New().Struct(result); err != nil {
		return result, err
	}
	return result, nil
}
