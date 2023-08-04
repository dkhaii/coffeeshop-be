package entity

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
