package config

import "time"

type (
	Global struct {
		Logger      *Logger      `mapstructure:"logger"`
		Listen      *Listen      `mapstructure:"listen"`
		DB          *DB          `mapstructure:"db"`
		Application *Application `mapstructure:"application"`
	}

	Logger struct {
	}

	Listen struct {
		HTTPServer *HTTPServer `mapstructure:"http"`
	}

	HTTPServer struct {
		Port                     uint16        `mapstructure:"port"`
		MaxRequestProcessingTime time.Duration `mapstructure:"maxRequestProcessingTime"`
	}

	DB struct {
		Host              string `mapstructure:"host"`
		Port              uint16 `mapstructure:"port"`
		User              string `mapstructure:"user"`
		Password          string `mapstructure:"password"`
		Name              string `mapstructure:"name"`
		SSLMode           string `mapstructure:"sslMode"`
		Schema            string `mapstructure:"schema"`
		ApplicationName   string `mapstructure:"applicationName"`
		MigrationsDirPath string `mapstructure:"migrationsDirPath"`
	}

	Application struct {
		PasswordHash string `mapstructure:"passwordHash"`
	}
)

func NewConfig() *Global {
	return &Global{}
}
