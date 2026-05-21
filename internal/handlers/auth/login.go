package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aamedvedevTECH/chat-backend/internal/forms/auth"
)

type LoginByPassword struct {
}

func NewLoginByPasswordHandler() *LoginByPassword {
	return &LoginByPassword{}
}

func (l *LoginByPassword) Call(w http.ResponseWriter, r *http.Request) {
	form := new(auth.LoginByPassword)

	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println(form.Username, form.Password)
}
