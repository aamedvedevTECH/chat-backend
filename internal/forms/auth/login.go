package auth

type LoginByPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
