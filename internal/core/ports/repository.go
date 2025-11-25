package ports

import "github.com/doble97/scheduleapi/internal/core/domain"

type DashboardRepository interface {
	Save(domain.Dashboard) (domain.Dashboard, error)
	FindByIDUserOne(id int) (domain.Dashboard, error)
	FindByIDUserMany(id int) ([]domain.Dashboard, error)
	Update(domain.Dashboard) (domain.Dashboard, error)
	Delete(id string) error
}
