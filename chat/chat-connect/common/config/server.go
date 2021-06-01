package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Server struct {
	Profile      string
	ProtocolType string
	IP           string
	Port         uint16
	Network      string
}

func InitServer(cfg *viper.Viper) *Server {
	return &Server{
		Profile:      strings.ToUpper(cfg.GetString("profile")),
		ProtocolType: strings.ToUpper(cfg.GetString("protocolType")),
		IP:           cfg.GetString("ip"),
		Port:         uint16(cfg.GetUint("port")),
		Network:      cfg.GetString("network"),
	}
}

var ServerConfig *Server = nil
