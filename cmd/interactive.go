package cmd

import (
	"bufio"
	"fmt"
	"go-todo-app/internal/todo"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// interactiveCmd represents the interactive command
var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Start interactive mode",
	Run: func(cmd *cobra.Command, args []string) {
		runInteractiveMode()
	},
}

func init() {
	rootCmd.AddCommand(interactiveCmd)
}

func printBanner() {
	banner := `
   _____       _______        _       
  / ____|     |__   __|      | |      
 | |  __  ___    | | ___   __| | ___  
 | | |_ |/ _ \   | |/ _ \ / _  |/ _ \ 
 | |__| | (_) |  | | (_) | (_| | (_) |
  \_____|\___/   |_|\___/ \__,_|\___/ 
                                      
  Interactive Todo Manager                              
`
	fmt.Println(banner)
}

func runInteractiveMode() {
	scanner := bufio.NewScanner(os.Stdin)
	
	printBanner()
	fmt.Println("--------------------------------")
	printHelp()
	
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}
		
		input := scanner.Text()
		args := strings.Fields(input)
		
		if len(args) == 0 {
			continue
		}
		
		command := args[0]
		cmdArgs := args[1:]
		
		switch command {
		case "add":
			handleAdd(cmdArgs)
		case "list":
			handleList(cmdArgs)
		case "done":
			handleDone(cmdArgs)
		case "rm":
			handleRemove(cmdArgs)
		case "help":
			printHelp()
		case "quit", "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  add [title]      - Add a new todo")
	fmt.Println("  list             - List all todos")
	fmt.Println("  list -a          - List all todos including completed ones")
	fmt.Println("  done [id]        - Mark a todo as done")
	fmt.Println("  rm [id]          - Remove a todo")
	fmt.Println("  help             - Show this help menu")
	fmt.Println("  quit, exit       - Exit the program")
}

func handleAdd(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a title for the todo")
		return
	}
	
	title := strings.Join(args, " ")
	newTodo := todo.AddTodo(title)
	
	if err := todo.SaveToFile(); err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	
	fmt.Printf("Added todo: %s (ID: %d)\n", newTodo.Title, newTodo.ID)
}

func handleList(args []string) {
	showCompleted := false
	for _, arg := range args {
		if arg == "-a" || arg == "--all" {
			showCompleted = true
			break
		}
	}
	
	todos := todo.GetAllTodos()
	if len(todos) == 0 {
		fmt.Println("No todos found!")
		return
	}
	
	fmt.Println("Your todos:")
	fmt.Println("------------------")
	
	for _, t := range todos {
		if t.Completed && !showCompleted {
			continue
		}
		
		status := " "
		if t.Completed {
			status = "✓"
		}
		
		fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
	}
}

func handleDone(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide the ID of the todo to mark as done")
		return
	}
	
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
}

func handleRemove(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide the ID of the todo to remove")
		return
	}
	
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
}