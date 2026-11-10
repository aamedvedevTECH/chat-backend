package users

import "github.com/aamedvedevTECH/chat-backend/internal/models/users"

// TODO: может переименовать??
type CreateUserRepo struct {
}

func NewCreateUserRepo() *CreateUserRepo {
	return &CreateUserRepo{}
}

func (r *CreateUserRepo) Call(userData users.RegisterData) {
	const sq = `
		INSERT INTO users (
			id,
			email,
			login,
			password_hash,
			name
		)
		VALUES ($1, $2, $3, $4, $5)
	`

}
