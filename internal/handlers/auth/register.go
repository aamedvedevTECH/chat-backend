package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aamedvedevTECH/chat-backend/internal/entities/users"
	"github.com/aamedvedevTECH/chat-backend/internal/forms/auth"
)

type Register struct {
	registerUser RegisterUserUseCase
}

func NewRegisterHandler(registerUser RegisterUserUseCase) *Register {
	return &Register{
		registerUser: registerUser,
	}
}

func (u *Register) Call(w http.ResponseWriter, r *http.Request) {
	form := new(auth.RegisterForm)

	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println(form.Username, form.Password)

	// TODO: форма validate required

	user, tokenData, err := u.registerUser.Call(&users.RegisterData{
		Name:     form.Name,
		Email:    form.Email,
		Username: form.Username,
		Password: form.Password,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return http.
}
