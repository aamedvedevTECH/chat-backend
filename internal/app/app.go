package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"

	"github.com/aamedvedevTECH/chat-backend/config"
	"github.com/aamedvedevTECH/chat-backend/internal/handlers/auth"
	"github.com/aamedvedevTECH/chat-backend/internal/router"
)

type App struct {
	logger     *zap.SugaredLogger
	config     *config.Global
	httpServer *http.Server
	db         *sqlx.DB
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

	if err = a.initDB(ctx); err != nil {
		return fmt.Errorf("creating db: %w", err)
	}

	if err = a.migrate(ctx); err != nil {
		return fmt.Errorf("applying migrations: %w", err)
	}

	a.httpServer = router.Router(
		a.config,
		loginByPassword,
	)

	return nil
}

func (a *App) initDB(ctx context.Context) error {
	dataSourceName := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s search_path=%s application_name=%s",
		a.config.DB.User,
		a.config.DB.Password,
		a.config.DB.Host,
		a.config.DB.Port,
		a.config.DB.Name,
		a.config.DB.SSLMode,
		a.config.DB.Schema,
		a.config.DB.ApplicationName,
	)

	connConfig, err := pgx.ParseConfig(dataSourceName)
	if err != nil {
		return fmt.Errorf("failed to parse database connection string: %w", err)
	}

	pgxdb := stdlib.OpenDB(*connConfig)

	a.db = sqlx.NewDb(pgxdb, "pgx")
	if err = a.db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	a.logger.Info("database connection established")

	return nil
}

func (a *App) migrate(ctx context.Context) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("goose.SetDialect err: %s", err)
	}

	if err := goose.UpContext(ctx, a.db.DB, a.config.DB.MigrationsDirPath); err != nil {
		return fmt.Errorf("goose.Up err: %s", err)
	}

	return nil
}

func (a *App) Close(ctx context.Context) {
	// Trigger graceful shutdown
	if err := a.httpServer.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	if a.db != nil {
		if err := a.db.Close(); err != nil {
			log.Printf("database close err: %s", err)
		}
	}
}
