package todo

import (
	"time"
)

// Todo represents a single todo item
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}

var todoList TodoList

func GetNextId() int {
	if len(todoList.Todos)==0 {
		return 1
	}
	return todoList.Todos[len(todoList.Todos)-1].ID + 1
}

func AddTodo(title string) Todo {
	todo := Todo{
		ID:        GetNextId(),
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	todoList.Todos = append(todoList.Todos, todo)
	return todo
}

func GetAllTodos() []Todo {
	return todoList.Todos
}

func MarkAsDone(id int) (bool, Todo) {
	for i := range todoList.Todos {
		if todoList.Todos[i].ID == id {
			todoList.Todos[i].Completed = true
			return true, todoList.Todos[i]
		}
	}
	return false, Todo{}
}

// RemoveTodo removes a todo by ID
func RemoveTodo(id int) (bool, Todo) {
	for i, todo := range todoList.Todos {
		if todo.ID == id {
			// Save the todo to return it later
			removedTodo := todo
			
			// Remove the todo by slicing it out
			todoList.Todos = append(todoList.Todos[:i], todoList.Todos[i+1:]...)
			return true, removedTodo
		}
	}
	return false, Todo{}
}