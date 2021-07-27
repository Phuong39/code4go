package napodate

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewService(t *testing.T) {
	service := NewService()
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(service.Status(ctx))
	fmt.Println(service.Get(ctx))
}
