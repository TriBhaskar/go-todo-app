// main.go
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "myapp",
		Short: "My CLI application",
		Long:  `A longer description of my awesome CLI application built with Go and Cobra.`,
		Run: func(cmd *cobra.Command, args []string) {
			// This is the code that will be executed when the root command is called
			fmt.Println("Welcome to my CLI app! Use --help to see available commands.")
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `Print the version number of the CLI application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("myapp v0.1 - Alpha")
		},
	}

	var getCmd = &cobra.Command{
		Use:   "get [resource]",
		Short: "Get a resource",
		Long:  `Get a specific resource by name`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Getting resource: %s\n", args[0])
			
			// Access flags
			all, _ := cmd.Flags().GetBool("all")
			if all {
				fmt.Println("Fetching all related resources too")
			}
		},
	}

	// Add flags to commands
	getCmd.Flags().BoolP("all", "a", false, "Get all resources")
	
	// Add commands to root
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(getCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}