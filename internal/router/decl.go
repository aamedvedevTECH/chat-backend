package router

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/aamedvedevTECH/chat-backend/config"
)

func Router(
	cfg *config.Global,
	loginByPassword LoginByPasswordHandler,
) *http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/v1/auth/login", loginByPassword.Call)

	server := &http.Server{Addr: fmt.Sprint(":", cfg.Listen.HTTPServer.Port), Handler: r}

	// Run server in the background
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	return server
}
