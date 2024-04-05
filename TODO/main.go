
package main

import (
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	ID          int
	Description string
	Completed   bool
}

var todos []*Todo
var idCounter = 1

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [command]")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add [description]")
			os.Exit(1)
		}
		description := os.Args[2]
		todo := AddTodo(description)
		fmt.Printf("Added todo with ID %d\n", todo.ID)
	case "list":
		ListTodos()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo complete [ID]")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			os.Exit(1)
		}
		err = CompleteTodo(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Marked todo with ID %d as completed\n", id)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete [ID]")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			os.Exit(1)
		}
		err = DeleteTodo(id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Deleted todo with ID %d\n", id)
	default:
		fmt.Println("Invalid command. Usage: todo [command]")
		os.Exit(1)
	}
}

func AddTodo(description string) *Todo {
	todo := &Todo{
		ID:          idCounter,
		Description: description,
		Completed:   false,
	}
	todos = append(todos, todo)
	idCounter++
	return todo
}

func ListTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos found")
		return
	}
	for _, todo := range todos {
		status := "Incomplete"
		if todo.Completed {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Description: %s, Status: %s\n", todo.ID, todo.Description, status)
	}
}

func CompleteTodo(id int) error {
	for _, todo := range todos {
		if todo.ID == id {
			todo.Completed = true
			return nil
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}

func DeleteTodo(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}
