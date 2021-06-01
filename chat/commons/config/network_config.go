package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Network struct {
	ProtocolType string
	IP           string
	Port         uint16
	Network      string
}

func InitServer(cfg *viper.Viper) *Network {
	return &Network{
		ProtocolType: strings.ToUpper(cfg.GetString("protocolType")),
		IP:           cfg.GetString("ip"),
		Port:         uint16(cfg.GetUint("port")),
		Network:      cfg.GetString("network"),
	}
}

var NetworkConfig *Network = nil
