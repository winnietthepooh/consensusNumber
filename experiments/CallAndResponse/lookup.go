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

	for i, srv := range list {
		fmt.Printf("%d: %s %s\n", i, srv.Type, srv.Location)
	}
}
