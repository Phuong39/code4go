package config

import "github.com/spf13/viper"

type Logger struct {
	Path  string
	Level string
}

func InitLogger(cfg *viper.Viper) *Logger {
	logger := &Logger{
		Path:  cfg.GetString("Path"),
		Level: cfg.GetString("Level"),
	}
	return logger
}

var LoggerConfig *Logger = nil
