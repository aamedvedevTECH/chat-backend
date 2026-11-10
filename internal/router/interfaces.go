package router

import "net/http"

type LoginByPasswordHandler interface {
	Call(w http.ResponseWriter, r *http.Request)
}
