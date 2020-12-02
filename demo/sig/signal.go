package main

import (
	"fmt"
	log "github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		log.Infof("goim-comet get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("退出")
			return
		case syscall.SIGHUP:
			fmt.Println("SIGHUP")
		default:
			fmt.Println("default")
			return
		}
	}
}
