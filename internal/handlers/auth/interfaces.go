package auth

import (
	"github.com/aamedvedevTECH/chat-backend/internal/entities/auth"
	"github.com/aamedvedevTECH/chat-backend/internal/entities/users"
)

type RegisterUserUseCase interface {
	Call(registerData *users.RegisterData) (*users.User, *auth.TokenData, error)
}
