package users

import (
	"github.com/google/uuid"

	"github.com/aamedvedevTECH/chat-backend/config"
	"github.com/aamedvedevTECH/chat-backend/internal/entities/auth"
	"github.com/aamedvedevTECH/chat-backend/internal/entities/users"
	users_model "github.com/aamedvedevTECH/chat-backend/internal/models/users"
	"github.com/aamedvedevTECH/chat-backend/pkg/helper"
)

type RegisterUser struct {
	applicationConfig *config.Application
}

func NewRegisterUserUseCase(applicationConfig *config.Application) *RegisterUser {
	return &RegisterUser{
		applicationConfig: applicationConfig,
	}
}

func (r *RegisterUser) Call(registerData *users.RegisterData) (*users.User, *auth.TokenData, error) {
	passwordHash := helper.PasswordHash(registerData.Password, r.applicationConfig.PasswordHash)

	data := users_model.RegisterData{
		ID:           uuid.New(),
		Email:        registerData.Email,
		Username:     registerData.Username,
		PasswordHash: passwordHash,
		Name:         registerData.Name,
	}

	// TODO: create auth session
	// TODO: вшиваем в токен userID / refreshID и ????? валидность точно exp / issuedAt тоже бы
	// id (uuid.New()) / user_id (уже будет с прошлого шага), valid_after_at / valid_before_at / state (в каком статусе сессия)
}
