package cmd

import (
	"fmt"
	"go-todo-app/internal/todo"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new todo",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		for i := 1; i < len(args); i++ {
			title += " " + args[i]
		}

		newTodo := todo.AddTodo(title)
		
		if err := todo.SaveToFile(); err != nil {
			fmt.Println("Error saving todo:", err)
			return
		}
		
		fmt.Printf("Added todo: %s (ID: %d)\n", newTodo.Title, newTodo.ID)
	},
}