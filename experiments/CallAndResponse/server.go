package main

import (
	"github.com/grandcat/zeroconf"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server, err := zeroconf.Register("test", "test.tcp", "local.", 8080, []string{"test"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
