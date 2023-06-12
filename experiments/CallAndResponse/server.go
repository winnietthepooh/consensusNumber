package main

import (
	"github.com/oleksandr/bonjour"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server, err := bonjour.Register("Test Service", "_test._tcp", "", 8080, []string{"txtv=1", "app=test"}, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Shutdown()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
