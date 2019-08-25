package service

import (
	"log"

	"github.com/spf13/cobra"
)

func startService(cmd *cobra.Command, args []string) {
	s, err := newService(cmd)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func stopService(cmd *cobra.Command, args []string) {
	s, err := newService(cmd)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Stop()
	if err != nil {
		log.Fatal(err)
	}
}
