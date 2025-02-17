package readcommands

import (
	"fmt"
	Todo "todo/todo"
)

func ReadCommands(command, msg, title string) {

	switch command {
	case "add":
		Todo.AddTodo(msg, title)
	case "rm":
		Todo.Remove(msg)
	case "lt":
		Todo.ListTodo()
	default:
		fmt.Println("invalid command")
	}

}
