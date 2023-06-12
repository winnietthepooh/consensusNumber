package main

import (
	"github.com/koron/go-ssdp"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	rn := rand.Intn(30)
	rns := strconv.Itoa(rn)
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	ad, err := ssdp.Advertise(
		"my:device",
		hostname+":"+rns,
		GetLocalIP(),
		"go-sddp sample",
		1800)
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	at := time.Tick(300 * time.Second)

loop:
	for {
		select {
		case <-sig:
			break loop
		case <-at:
			ad.Alive()
		}
	}

	ad.Bye()
	ad.Close()
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
