package main

import (
	"fmt"
	"os"

	"go-cli-tasks/internal/task"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task [add|list|complete]")
		return
	}

	command := os.Args[1]

	service := task.NewService()

	switch command {
	case "add":
		fmt.Println("add command received")
		// service.Add(...)
	case "list":
		fmt.Println("list command received")
		// service.List()
	case "complete":
		fmt.Println("complete command received")
		// service.Complete(...)
	default:
		fmt.Println("unknown command:", command)
	}
}
