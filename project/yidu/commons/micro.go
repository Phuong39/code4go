package common

import "github.com/micro/go-micro/v2"

var options []micro.Option

type MicroConfig struct {
	Version      string
	Name         string
	RegistryType string
}

func InitService(options ...micro.Option) {

}
