package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/encoding/charmap"
)

func sendExecToServer(cmd string, addr string, config *ssh.ClientConfig) string {

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

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	err = session.Run(cmd)

	outStr := "<Unreadable output>"
	d := charmap.CodePage850.NewDecoder()
	var status string

	if err != nil {
		fmt.Printf("Error during running command : %v\n", err)
		outStr, err = d.String(stderrBuf.String())
		status = Red("Failed").String()
	} else {
		outStr, err = d.String(stdoutBuf.String())
		status = Green("Succeed").String()
	}

	// bug : it seems the last line of output is the pwd for the next command (stdin)
	// need to remove it for a cleaner output
	idx := strings.LastIndex(outStr, "\n")
	outStr = outStr[:idx]

	return fmt.Sprintf("%s : %s\nOutput:\n%s\n", addr, status, outStr)
}

func checkHostKey(host string, remote net.Addr, key ssh.PublicKey) error {

	homeDir := os.Getenv("HOME")

	if runtime.GOOS == "windows" {
		homeDir = os.Getenv("USERPROFILE")
	}

	file, err := os.Open(filepath.Join(homeDir, ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		} else {

		}
	}

	if hostKey == nil {
		log.Printf("Warning : The hostkey wasn't found for %s", host)
	}

	// TODO : Handle errors, or ask user to add the public key to known hosts

	return nil
}

func execOnServers(cmd *cobra.Command, args []string) {

	var username string

	if runtime.GOOS == "windows" {
		username = os.Getenv("USERNAME")
	} else {
		username = os.Getenv("LOGNAME")
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password("123"),
			//ssh.PublicKeys(signer),
		},
		HostKeyCallback: checkHostKey,
	}

	/*
		key, err := ioutil.ReadFile("id_rsa")
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			log.Fatalf("unable to parse private key: %v", err)
		}*/

	// TODO : Get the list of servers using disco

	hosts := []string{"127.0.0.1:22000"}

	exec, _ := cmd.Flags().GetString("command")

	results := make(chan string, 10)
	timeout := time.After(5 * time.Second)

	for _, host := range hosts {
		go func(h string) {
			results <- sendExecToServer(exec, h, config)
		}(host)
	}

	for i := 0; i < len(hosts); i++ {
		select {
		case res := <-results:
			fmt.Println(res)
		case <-timeout:
			return
		}
	}
}
