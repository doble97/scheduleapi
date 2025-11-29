package error_app

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

func MapErrorToHTTP(err error) ErrorMapper {
	switch {
	case errors.Is(err, domain.ErrInvalidData):
		return InvalidInputError
	case errors.Is(err, domain.ErrAlreadyExists):
		return AlreadyExistsError
	case errors.As(err, new(validator.ValidationErrors)):
		return InvalidInputError // 422 Unprocessable entity

	case errors.Is(err, io.EOF), errors.As(err, new(*json.SyntaxError)),
		errors.As(err, new(*json.UnmarshalTypeError)):
		return BadRequestError
	default:
		return InternalServerError // 500 Internal server error
	}
}
