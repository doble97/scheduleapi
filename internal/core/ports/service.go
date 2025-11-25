package ports

import "github.com/doble97/scheduleapi/internal/core/domain"

// 🔑 DIRECTIVA PARA GENERAR EL MOCK DEL SERVICIO
//
//go:generate mockgen -destination=../../platform/api/http/mocks/mock_dashboard_service.go -package=mocks github.com/doble97/scheduleapi/internal/core/ports DashboardService
type DashboardService interface {
	CreateDashboard(domain.Dashboard) (domain.Dashboard, error)
	GetOneDashboardByIDUser(id int) (domain.Dashboard, error)
	GetManyDashboardsByIDUser(id int) ([]domain.Dashboard, error)
	UpdateDashboard(domain.Dashboard) (domain.Dashboard, error)
	DeleteDashboard(int) error
}
