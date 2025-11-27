package mariadb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/core/ports"
)

type UserRepoMariaDB struct {
	conn *sql.DB
}

// GetUserByEmail implements ports.UserRepository.
func (u *UserRepoMariaDB) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user := domain.User{}
	query := "SELECT id, name, last_name, email, password, created_at FROM users WHERE email = ?"
	err := u.conn.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

// SaveUser implements ports.UserRepository.
func (u *UserRepoMariaDB) SaveUser(ctx context.Context, user domain.User) (int, error) {
	query := `INSERT INTO users (name, last_name, email, password) VALUES (?, ?, ?, ?)`
	result, err := u.conn.ExecContext(ctx, query, user.Name, user.LastName, user.Email, user.Password)
	if err != nil {
		return 0, nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func NewUserRepoMariaDB(dbConn *sql.DB) ports.UserRepository {
	return &UserRepoMariaDB{
		conn: dbConn,
	}
}
