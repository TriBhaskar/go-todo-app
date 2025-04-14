package cmd

import (
	"fmt"
	"go-todo-app/internal/todo"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "rm [id]",
	Short: "Remove a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a number.")
			return
		}
		
		found, _ := todo.RemoveTodo(id)
		if !found {
			fmt.Printf("Todo with ID %d not found\n", id)
			return
		}
		
		if err := todo.SaveToFile(); err != nil {
			fmt.Println("Error saving todos:", err)
			return
		}
		
		fmt.Printf("Removed todo %d\n", id)
	},
}