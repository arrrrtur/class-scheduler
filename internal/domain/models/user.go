package models

import "time"

type User struct {
	ID        string
	Username  string
	Group     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
