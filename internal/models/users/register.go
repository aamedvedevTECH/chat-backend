package users

import "github.com/google/uuid"

type RegisterData struct {
	ID           uuid.UUID
	Email        string
	Username     string
	PasswordHash string
	Name         string
}
