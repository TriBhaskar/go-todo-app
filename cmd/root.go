package cmd

import (
	"fmt"
	"go-todo-app/internal/todo"
	"os"

	"github.com/spf13/cobra"
)

var (
	todoFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLI task manager",
	Long:  `A simple CLI Todo application built with Go and Cobra that helps you manage your tasks.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Load todos before any command runs
		err := todo.LoadFromFile(todoFile)
		if err != nil {
			fmt.Println("Error loading todos:", err)
			os.Exit(1)
		}
	},
	// This will run if no subcommand is provided
	Run: func(cmd *cobra.Command, args []string) {
		// Start interactive mode by default
		runInteractiveMode()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Set default todo file location
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}
	todoFile = homeDir + "/.todos.json"
	todo.SetTodoFile(todoFile)

	// Add commands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(removeCmd)
}