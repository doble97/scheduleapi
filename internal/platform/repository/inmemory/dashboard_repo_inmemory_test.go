package inmemory

import (
    "testing"

    "github.com/doble97/scheduleapi/internal/core/domain"
)

// Red test: require non-empty Title on Save. Current implementation does not
// enforce this, so the test will fail until validation is added (TDD - red).
func TestDashboardRepo_Save_RequiresTitle(t *testing.T) {
    repo := NewInMemoryDashboardRepo()

    _, err := repo.Save(domain.Dashboard{Title: ""})
    if err == nil {
        t.Fatalf("expected error when saving dashboard without title, got nil")
    }
}
