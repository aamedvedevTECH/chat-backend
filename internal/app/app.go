package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aamedvedevTECH/chat-backend/internal/handlers/auth"
	"go.uber.org/zap"

	"github.com/aamedvedevTECH/chat-backend/config"
	"github.com/aamedvedevTECH/chat-backend/internal/router"
)

type App struct {
	logger     *zap.SugaredLogger
	config     *config.Global
	httpServer *http.Server
}

func NewApp(globalConfig *config.Global) *App {
	return &App{
		config: globalConfig,
	}
}

func (a *App) Start(ctx context.Context) error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("creating logger: %w", err)
	}

	a.logger = logger.Sugar()

	// handlers
	var (
		loginByPassword = auth.NewLoginByPasswordHandler()
	)

	a.httpServer = router.Router(
		a.config,
		loginByPassword,
	)

	return nil
}

func (a *App) Close(ctx context.Context) {
	// Trigger graceful shutdown
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
