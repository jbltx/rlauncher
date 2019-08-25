package service

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func installService(cmd *cobra.Command, args []string) {

	s, err := newService(cmd)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Install()
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "denied") {
			msg += "\nYou need to execute this command as administrator (or root)"
		}
		log.Fatal(msg)
	}
}

func uninstallService(cmd *cobra.Command, args []string) {

	s, err := newService(cmd)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Uninstall()
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "denied") {
			msg += "\nYou need to execute this command as administrator (or root)"
		}
		log.Fatal(msg)
	}
}
