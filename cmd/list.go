package cmd

import (
	"fmt"
	"go-todo-app/internal/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		showCompleted, _ := cmd.Flags().GetBool("all")
		
		todos := todo.GetAllTodos()
		if len(todos) == 0 {
			fmt.Println("No todos found!")
			return
		}
		
		fmt.Println("Your todos:")
		fmt.Println("------------------")
		
		for _, todo := range todos {
			if todo.Completed && !showCompleted {
				continue
			}
			
			status := " "
			if todo.Completed {
				status = "✓"
			}
			
			fmt.Printf("[%s] %d: %s\n", status, todo.ID, todo.Title)
		}
	},
}

func init() {
	// Add flags to list command
	listCmd.Flags().BoolP("all", "a", false, "Show completed todos as well")
}