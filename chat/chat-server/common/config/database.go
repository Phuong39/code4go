package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Database struct {
	Driver string
	Source string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Driver: strings.ToUpper(cfg.GetString("driver")),
		Source: strings.ToUpper(cfg.GetString("source")),
	}
}

var DatabaseConfig *Database = nil
