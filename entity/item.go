package entity

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID
	Name      string
	Stock     int
	Price     int
	Img       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
