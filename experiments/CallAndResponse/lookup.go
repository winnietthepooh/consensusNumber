package main

import (
	"fmt"
	"github.com/hashicorp/mdns"
)

func main() {

	entriesCh := make(chan *mdns.ServiceEntry)
	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	mdns.Lookup("test.tcp", entriesCh)
	close(entriesCh)

}
