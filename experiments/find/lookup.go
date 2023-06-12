package main

import (
	"fmt"
	"github.com/koron/go-ssdp"
)

func main() {
	list, err := ssdp.Search("my:device", 1, "")
	if err != nil {
		panic(err)
	}
	ips := []string{}
	for i, srv := range list {
		if len(ips) > 0 {
			if ips[len(ips)-1] == srv.Location {
				continue
			}
		}

		fmt.Printf("%d: %s %s\n", i, srv.Type, srv.Location)
		ips = append(ips, srv.Location)
	}
}
