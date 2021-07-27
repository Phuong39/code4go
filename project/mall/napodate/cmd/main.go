package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"mall/napodate"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our napodate service
	srv := napodate.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// 映射端点
	endpoints := napodate.Endpoints{
		GetEndpoint:      napodate.MakeGetEndpoint(srv),
		StatusEndpoint:   napodate.MakeStatusEndpoint(srv),
		ValidateEndpoint: napodate.MakeValidateEndpoint(srv),
	}

	// HTTP 传输
	go func() {
		gin.SetMode(gin.ReleaseMode)
		log.Println("napodate is listening on port:", *httpAddr)
		handler := napodate.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
