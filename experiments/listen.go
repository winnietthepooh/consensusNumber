package main

import (
	"github.com/pion/mdns"
	"golang.org/x/net/ipv4"
	"net"
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

	_, err = mdns.Server(ipv4.NewPacketConn(listener), &mdns.Config{
		LocalNames: []string{"test.local"},
	})

	if err != nil {
		panic(err)
	}
	select {}
}
