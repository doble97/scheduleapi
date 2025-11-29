package mariadb

import (
	"context"
	"database/sql"
	"time"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
	_ "github.com/go-sql-driver/mysql"
)

type DashboardRepoMariaDB struct {
	// Add necessary fields for database connection
	conn *sql.DB
}

// Delete implements ports.DashboardRepository.
func (d *DashboardRepoMariaDB) Delete(id string) error {
	panic("unimplemented")
}

// FindByIDUserMany implements ports.DashboardRepository.
func (d *DashboardRepoMariaDB) FindByIDUserMany(id int) ([]domain.Dashboard, error) {
	// panic("unimplemented")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := d.conn.QueryContext(ctx, "SELECT * FROM dashboards")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dashboards := make([]domain.Dashboard, 0)
	for rows.Next() {
		var dash domain.Dashboard
		err := rows.Scan(&dash.ID, &dash.Title, &dash.Description)
		if err != nil {
			return nil, err
		}
		dashboards = append(dashboards, dash)
	}
	return dashboards, nil
}

// FindByIDUserOne implements ports.DashboardRepository.
func (d *DashboardRepoMariaDB) FindByIDUserOne(id int) (domain.Dashboard, error) {
	panic("unimplemented")
}

// Save implements ports.DashboardRepository.
func (d *DashboardRepoMariaDB) Save(dashboard domain.Dashboard) (domain.Dashboard, error) {
	// panic("unimplemented")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "INSERT INTO dashboards (title, description) VALUES (?, ?)"
	result, err := d.conn.ExecContext(ctx, query, dashboard.Title, dashboard.Description)
	if err != nil {
		return domain.Dashboard{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return domain.Dashboard{}, err
	}
	dashboard.ID = int(id)
	return dashboard, nil
}

// Update implements ports.DashboardRepository.
func (d *DashboardRepoMariaDB) Update(domain.Dashboard) (domain.Dashboard, error) {
	panic("unimplemented")
}

func NewDashboardRepoMariaDB(dbConn *sql.DB) ports.DashboardRepository {
	return &DashboardRepoMariaDB{
		conn: dbConn,
	}
}
