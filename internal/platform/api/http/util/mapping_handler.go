package util

import (
	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
)

func DomainToDTOResponse(dash domain.Dashboard) dto.DashboardResponse {
	return dto.DashboardResponse{ID: dash.ID, Title: dash.Title, Description: dash.Description}
}

func UserDomainToAuthResponse(user domain.User, token string) dto.AuthResponse {
	return dto.AuthResponse{Token: token, UserID: user.ID, Email: user.Email}
}
