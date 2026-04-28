package config

import "time"

type (
	Global struct {
		Logger *Logger `mapstructure:"logger"`
		Listen *Listen `mapstructure:"listen"`
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
)
