package service

import (
	"os"

	svc "github.com/kardianos/service"
	"github.com/spf13/cobra"
)

func newService(cmd *cobra.Command) (svc.Service, error) {

	port, _ := cmd.Flags().GetInt("port")
	nodisco, _ := cmd.Flags().GetBool("no-discovery")
	discoPort, _ := cmd.Flags().GetInt("discovery-port")

	svcConfig := &svc.Config{
		Name:        "rLauncherService",
		DisplayName: "rLauncher Service",
		Description: "Launch at startup a custom SSH server which can be handled by authorized rLauncher clients",
		Arguments:   []string{"service", "run"},
	}

	server := &rLauncherServer{
		Config: &serverConfig{
			Port:          port,
			DiscoveryPort: discoPort,
			UseDiscovery:  !nodisco,
		},
	}

	return svc.New(server, svcConfig)
}

// NewServiceCmd returns an instance of the Service sub command
func NewServiceCmd() *cobra.Command {

	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Service related commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	installCmd := &cobra.Command{
		Use:   "install",
		Short: "Install Service",
		Long: `
Install the application as a Service (Windows, Linux and OSX compatible)
		`,
		Run: installService,
	}

	uninstallCmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall Service",
		Long: `
Uninstall the application as a Service (Windows, Linux and OSX compatible)
		`,
		Run: uninstallService,
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the service",
		Run:   startService,
	}

	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop the service",
		Run:   stopService,
	}

	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the service",
		Run:   runService,
	}

	serviceCmd.AddCommand(installCmd)
	serviceCmd.AddCommand(uninstallCmd)
	serviceCmd.AddCommand(startCmd)
	serviceCmd.AddCommand(stopCmd)
	serviceCmd.AddCommand(runCmd)

	runCmd.Flags().IntP("port", "p", 22000, "The port to listen to")
	runCmd.Flags().Bool("no-discovery", false, "Do not install the discovery service")
	runCmd.Flags().Int("discovery-port", 42424, "The port used by the discovery service")

	return serviceCmd
}
