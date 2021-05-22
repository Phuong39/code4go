package config

import "github.com/BurntSushi/toml"

var C Config

type Config struct {
	Title    string
	DB       database `toml:"database"`
	Server   server
	Registry registry
}

type server struct {
	Name    string
	Version string
}

type registry struct {
	RegistryType string
	Addrs        []string
}

type database struct {
	Type        string
	Host        string
	Port        int32
	User        string
	Pass        string
	Name        string
	Charset     string
	Level       string
	Debug       bool
	Stdout      bool
	MaxIdle     int32
	MaxOpen     int32
	MaxLifetime int32
}

func init() {
	if _, err := toml.DecodeFile("./config/config.toml", &C); err != nil {
		panic(err)
	}
}
