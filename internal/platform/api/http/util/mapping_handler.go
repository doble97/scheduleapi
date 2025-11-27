package util

import (
	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
)

func DomainToDTOResponse(dash domain.Dashboard) dto.DashboardResponse {
	return dto.DashboardResponse{ID: dash.ID, Title: dash.Title, Description: dash.Description}
}
