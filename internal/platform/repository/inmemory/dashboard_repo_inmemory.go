package inmemory

import (
	"errors"
	"sync"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
)

type DashboardRepoInMemory struct {
	sync.Mutex
	store map[int]*domain.Dashboard
}

// Delete implements ports.DashboardRepository.
func (d *DashboardRepoInMemory) Delete(id string) error {
	panic("unimplemented")
}

// FindByIDUserMany implements ports.DashboardRepository.
func (d *DashboardRepoInMemory) FindByIDUserMany(id int) ([]domain.Dashboard, error) {
	d.Lock()
	defer d.Unlock()
	dashboards := make([]domain.Dashboard, 0)
	for _, dashboard := range d.store {
		// if dashboard.ID == id {
		// 	dashboards = append(dashboards, *dashboard)
		// }
		dashboards = append(dashboards, *dashboard)
	}

	return dashboards, nil
}

// FindByIDUserOne implements ports.DashboardRepository.
func (d *DashboardRepoInMemory) FindByIDUserOne(id int) (domain.Dashboard, error) {
	panic("unimplemented")
}

// Save implements ports.DashboardRepository.
func (d *DashboardRepoInMemory) Save(dashboard domain.Dashboard) (domain.Dashboard, error) {
	// panic("unimplemented")
	d.Lock()
	defer d.Unlock()
	if dashboard.Title == "" {
		return domain.Dashboard{}, errors.New("title is required")
	}
	if dashboard.ID == 0 {
		dashboard.ID = newID(d.store)
	}
	// store a copy to avoid external mutation
	dash := dashboard
	d.store[dash.ID] = &dash
	return dash, nil
}

// Update implements ports.DashboardRepository.
func (d *DashboardRepoInMemory) Update(domain.Dashboard) (domain.Dashboard, error) {
	panic("unimplemented")
}

func NewInMemoryDashboardRepo() ports.DashboardRepository {
	return &DashboardRepoInMemory{
		store: make(map[int]*domain.Dashboard),
	}
}
func newID(store map[int]*domain.Dashboard) int {
	return len(store) + 1
}

//	func NewInMemoryDashboardRepo() *DashboardRepoInMemory {
//		return &DashboardRepoInMemory{
//			store: make(map[string]*domain.Dashboard),
//		}
//	}
// func (r *DashboardRepoInMemory) Save(dashboard domain.Dashboard) (domain.Dashboard, error) {
// 	r.Lock()
// 	defer r.Unlock()
// 	if dashboard.Title == "" {
// 		return domain.Dashboard{}, errors.New("title is required")
// 	}
// 	if dashboard.ID == "" {
// 		dashboard.ID = newID()
// 	}
// 	// store a copy to avoid external mutation
// 	d := dashboard
// 	r.store[dashboard.ID] = &d
// 	return dashboard, nil
// }

// func (r *DashboardRepoInMemory) FindByIDUser(id string) (domain.Dashboard, error) {
// 	r.Lock()
// 	defer r.Unlock()
// 	if dashboard, exists := r.store[id]; exists {
// 		return *dashboard, nil
// 	}
// 	return domain.Dashboard{}, nil
// }

// func (r *DashboardRepoInMemory) Update(dashboard domain.Dashboard) (domain.Dashboard, error) {
// 	r.Lock()
// 	defer r.Unlock()
// 	r.store[dashboard.ID] = &dashboard
// 	return dashboard, nil
// }

// func (r *DashboardRepoInMemory) Delete(id string) error {
// 	r.Lock()
// 	defer r.Unlock()
// 	delete(r.store, id)
// 	return nil
// }
