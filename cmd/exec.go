package cmd

import (
	"fmt"

	"golang.org/x/crypto/ssh"

	"github.com/spf13/cobra"
)

func sendExecToServer(exec string, addr string, config *ssh.ClientConfig) {
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	out, err := session.CombinedOutput(exec)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}

func execOnServers(cmd *cobra.Command, args []string) {
	config := &ssh.ClientConfig{
		User: "toto",
		Auth: []ssh.AuthMethod{
			ssh.Password("123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// TODO : Get the list of servers using disco

	adresses := []string{"127.0.0.1:22000"}

	exec, _ := cmd.Flags().GetString("command")

	for _, addr := range adresses {
		sendExecToServer(exec+"\n", addr, config)
	}

}
