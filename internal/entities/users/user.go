package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Email     string
	Login     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
