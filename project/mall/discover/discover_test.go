package discover

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"testing"
)

func TestNewKitDiscoverClient(t *testing.T) {
	discoverClient, err := NewKitDiscoverClient("127.0.0.1", 123)
	if err != nil {
		panic(err)
	}
	serviceName := "demo"
	discoverClient.Register(serviceName, serviceName+"-"+uuid.NewV4().String(), "", "", 10000, nil, log.Default())
}
