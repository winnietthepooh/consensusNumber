package main

import (
	"fmt"
	"github.com/oleksandr/bonjour"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	resolver, err := bonjour.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	results := make(chan *bonjour.ServiceEntry)

	go func(results chan *bonjour.ServiceEntry, exitCh chan<- bool) {
		for e := range results {
			fmt.Printf("%s", e.AddrIPv4)
			exitCh <- true
		}
	}(results, resolver.Exit)

	err = resolver.Browse("_test._tcp", ".local", results)
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
