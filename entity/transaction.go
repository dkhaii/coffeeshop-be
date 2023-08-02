package entity

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID
	Name string
	Type int
	CreatedAt time.Time
	UpdatedAt time.Time
}