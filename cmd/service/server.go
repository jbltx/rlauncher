package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os/exec"
	"runtime"

	"github.com/grandcat/zeroconf"
	svc "github.com/kardianos/service"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

type serverConfig struct {
	Port         int
	UseDiscovery bool
}

type rLauncherServer struct {
	Config      *serverConfig
	Discovery   *zeroconf.Server
	SSHListener net.Listener
	SSHConfig   *ssh.ServerConfig
}

func (server *rLauncherServer) Start(s svc.Service) error {
	go server.run()
	return nil
}

func (server *rLauncherServer) run() {

	var err error

	if server.Config.UseDiscovery {
		server.Discovery, err = zeroconf.Register("rLauncher", "_workstation._tcp", "local.", server.Config.Port, []string{"txtv=0", "lo=1", "la=2"}, nil)
		if err != nil {
			log.Fatal(err)
		}
	}

	server.SSHConfig = &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			log.Printf("User : %s / %s", c.User(), pass)
			return nil, nil
		},
	}

	privateBytes, err := ioutil.ReadFile("id_rsa")
	if err != nil {
		log.Fatal("Failed to load private key (./id_rsa)")
	}
	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}
	server.SSHConfig.AddHostKey(private)

	server.SSHListener, err = net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", server.Config.Port))
	if err != nil {
		log.Fatalf("Failed to listen on %d (%s)", server.Config.Port, err)
	}

	for {
		tcpConn, err := server.SSHListener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection (%s)", err)
			continue
		}
		sshConn, chans, reqs, err := ssh.NewServerConn(tcpConn, server.SSHConfig)
		if err != nil {
			log.Printf("Failed to handshake (%s)", err)
			continue
		}
		log.Printf("New SSH connection from %s (%s)", sshConn.RemoteAddr(), sshConn.ClientVersion())
		go ssh.DiscardRequests(reqs)
		go handleChannels(chans)
	}
}

func handleChannels(chans <-chan ssh.NewChannel) {
	// Service the incoming Channel channel in go routine
	for newChannel := range chans {
		log.Printf("Channel type : %s", newChannel.ChannelType())

		if t := newChannel.ChannelType(); t != "session" {
			newChannel.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type: %s", t))
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Printf("Could not accept channel (%s)", err)
			continue
		}

		go func(in <-chan *ssh.Request) {
			for req := range in {
				log.Printf("-> Request Type : %s", req.Type)

				ok := false

				switch req.Type {
				case "exec":
					ok = true

					command := string(req.Payload[4 : req.Payload[3]+4])
					log.Printf("-> Request Exec Payload : %s", command)

					windowsCmd := []string{"/K", command}
					unixCmd := []string{"-c", command}

					var cmd *exec.Cmd
					// this is for Windows, check os env to get the right shell
					if runtime.GOOS == "windows" {
						cmd = exec.Command("cmd", windowsCmd...)
					} else {
						cmd = exec.Command("sh", unixCmd...)
					}
					cmd.Stdout = channel
					cmd.Stderr = channel

					err = cmd.Start()
					if err != nil {
						log.Fatalf("could not run command (%s)", err)
					}

					// teardown session
					go func() {
						_, err := cmd.Process.Wait()
						if err != nil {
							log.Printf("failed to exit bash (%s)", err)
						}
						channel.Close()
						log.Printf("session closed")
					}()
				}

				req.Reply(ok, nil)
			}
		}(requests)
	}
}

func (server *rLauncherServer) Stop(s svc.Service) error {
	if server.Config.UseDiscovery && server.Discovery != nil {
		server.Discovery.Shutdown()
	}

	server.SSHListener.Close()

	return nil
}

func runService(cmd *cobra.Command, args []string) {
	s, err := newService(cmd)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
