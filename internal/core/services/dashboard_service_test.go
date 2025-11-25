package services

import (
	"testing"

	"github.com/doble97/scheduleapi/internal/core/domain"
)

type fakeRepo struct{}

// FindByIDUserMany implements ports.DashboardRepository.
func (f *fakeRepo) FindByIDUserMany(id int) ([]domain.Dashboard, error) {
	panic("unimplemented")
}

// FindByIDUserOne implements ports.DashboardRepository.
func (f *fakeRepo) FindByIDUserOne(id int) (domain.Dashboard, error) {
	panic("unimplemented")
}

func (f *fakeRepo) Save(d domain.Dashboard) (domain.Dashboard, error) {
	d.ID = 1
	return d, nil
}

func (f *fakeRepo) FindByIDUser(id int) (domain.Dashboard, error) {
	return domain.Dashboard{}, nil
}

func (f *fakeRepo) Update(d domain.Dashboard) (domain.Dashboard, error) {
	return domain.Dashboard{}, nil
}

func (f *fakeRepo) Delete(id string) error {
	return nil
}

func TestDashboardService_CreateDashboard(t *testing.T) {
	repo := &fakeRepo{}
	svc := NewDashboardService(repo)

	input := domain.Dashboard{Title: "Test Title"}

	got, err := svc.CreateDashboard(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.ID == 0 {
		t.Fatalf("expected generated ID, got empty")
	}

	if got.Title != input.Title {
		t.Fatalf("expected title %q, got %q", input.Title, got.Title)
	}
}
