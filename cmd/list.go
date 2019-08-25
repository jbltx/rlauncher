package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/spf13/cobra"
)

func listServers(cmd *cobra.Command, args []string) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}
	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			fmt.Printf("-> Hostname : %s\n   IPv4 : %s\n   Port : %d", entry.HostName, entry.AddrIPv4, entry.Port)
		}
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))
	defer cancel()
	err = resolver.Browse(ctx, "_workstation._tcp", "local", entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	time.Sleep(1 * time.Second)
}
