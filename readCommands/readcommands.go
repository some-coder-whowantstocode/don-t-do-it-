package readcommands

import (
	"fmt"
	Todo "todo/todo"
)

func ReadCommands(command string, msg string) {

	switch command {
	case "add":
		Todo.AddTodo(msg)
	case "rm":
		Todo.Remove(msg)
	default:
		fmt.Println("invalid command")
	}

}
