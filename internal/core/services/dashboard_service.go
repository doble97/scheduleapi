package services

import (
	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
)

type dashboardService struct {
	repo ports.DashboardRepository
}

// GetManyDashboardsByIDUser implements ports.DashboardService.
func (s *dashboardService) GetManyDashboardsByIDUser(id int) ([]domain.Dashboard, error) {
	return s.repo.FindByIDUserMany(id)
}

// GetOneDashboardByIDUser implements ports.DashboardService.
func (s *dashboardService) GetOneDashboardByIDUser(id int) (domain.Dashboard, error) {
	panic("unimplemented")
}

// CreateDashboard implements ports.DashboardService.
func (s *dashboardService) CreateDashboard(d domain.Dashboard) (domain.Dashboard, error) {
	return s.repo.Save(d)
}

// DeleteDashboard implements ports.DashboardService.
func (d *dashboardService) DeleteDashboard(id int) error {
	panic("unimplemented")
}

// GetDashboardByIDUser implements ports.DashboardService.
func (d *dashboardService) GetDashboardByIDUser(id int) (domain.Dashboard, error) {
	panic("unimplemented")
}

// UpdateDashboard implements ports.DashboardService.
func (d *dashboardService) UpdateDashboard(dashboard domain.Dashboard) (domain.Dashboard, error) {
	panic("unimplemented")
}

func NewDashboardService(repo ports.DashboardRepository) ports.DashboardService {
	return &dashboardService{
		repo: repo,
	}
}
