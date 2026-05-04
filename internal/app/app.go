package app

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"github.com/aamedvedevTECH/chat-backend/config"
)

type App struct {
	logger *zap.SugaredLogger
	config *config.Global
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

	return nil
}

func (a *App) Close(ctx context.Context) {

}
