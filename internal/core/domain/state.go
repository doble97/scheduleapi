package domain

import "time"

type State struct {
	ID        int
	UserId    int
	Name      string
	CreatedAt time.Time
}
