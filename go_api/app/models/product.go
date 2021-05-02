package models

import "time"

type Product struct {
	ID   int
	UUID string
	Name string
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
}
