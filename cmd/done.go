package cmd

import (
	"fmt"
	"go-todo-app/internal/todo"
	"strconv"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a todo as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID. Please provide a number.")
			return
		}
		
		found, _ := todo.MarkAsDone(id)
		if !found {
			fmt.Printf("Todo with ID %d not found\n", id)
			return
		}
		
		if err := todo.SaveToFile(); err != nil {
			fmt.Println("Error saving todo:", err)
			return
		}
		
		fmt.Printf("Marked todo %d as done\n", id)
	},
}