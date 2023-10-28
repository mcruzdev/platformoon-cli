package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mcruzdev/platformoon-cli/internal/application"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := NewPlatformoonRoot()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NewPlatformoonRoot() *cobra.Command {
	var root = &cobra.Command{
		Use:   "platformoon",
		Short: "Platformoon CLI facilitates programmers lifes",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("platformoon works")
		},
	}

	root.AddCommand(NewApplicationCmd())

	return root
}

func NewApplicationCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "application",
		Short: "This command consists of multiple subcommands to interact with applications",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("platformoon application works")
		},
	}

	cmd.AddCommand(NewApplicationCreateCmd())
	return cmd
}

func NewApplicationCreateCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "This command is used to create a new application",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("platformoon application create works")
			application.CreateNewApplication("application", "API", "JAVA")
		},
	}

	return cmd
}
