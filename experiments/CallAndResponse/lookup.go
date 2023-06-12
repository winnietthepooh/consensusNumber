package main

import (
	"context"
	"fmt"
	"github.com/grandcat/zeroconf"
	"time"
)

func main() {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		panic(err)
	}
	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			fmt.Println(entry.AddrIPv4[len(entry.AddrIPv4)-1])
		}
	}(entries)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = resolver.Browse(ctx, "test.tcp", "local.", entries)
	if err != nil {
		panic(err)
	}
	<-ctx.Done()
}
