package main

import (
	"context"
	"fmt"
	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", mdns.DefaultAddress)
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic(err)
	}

	server, err := mdns.Server(ipv4.NewPacketConn(listener), &mdns.Config{})

	if err != nil {
		panic(err)
	}

	answer, src, err := server.Query(context.TODO(), "test.local")
	if err != nil {
		panic(err)
	}
	fmt.Println(answer)
	fmt.Println(src)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
