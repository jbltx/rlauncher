package cmd

import (
	"fmt"
	"os"

	"github.com/jbltx/rlauncher/cmd/service"
	"github.com/spf13/cobra"
)

// Execute is the main entrypoint of the CLI app
func Execute() {

	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(0)
		},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List servers",
		Long:  "List all available servers on the current local network",
		Run:   listServers,
	}

	execCmd := &cobra.Command{
		Use:   "exec",
		Short: "Execute command",
		Long:  "Execute a command on given servers",
		Run:   execOnServers,
	}

	// TODO : Add filters as arguments in the EXEC command line
	// execCmd.Flags().Int()
	execCmd.Flags().StringP("command", "c", "echo hello world !", "The command to execute remotely")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(execCmd)

	rootCmd.AddCommand(service.NewServiceCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
