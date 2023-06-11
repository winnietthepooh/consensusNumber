package main

import (
	"github.com/hashicorp/mdns"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {
	num := rand.Int()
	snum := strconv.Itoa(num)
	info := []string{"service"}
	service, _ := mdns.NewMDNSService(snum, "test.tcp", "", "", 8000, nil, info)
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
