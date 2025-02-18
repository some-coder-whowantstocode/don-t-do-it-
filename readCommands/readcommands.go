package readcommands

import (
	"fmt"
	Todo "todo/todo"
)

func ReadCommands(command, msg, title, id string) {

	switch command {
	case "add":
		Todo.AddTodo(msg, title)
	case "rm":
		Todo.Remove(msg)
	case "lt":
		arr := Todo.ListTask("./task.todo")
		Todo.ConsoleTodo(arr)
	case "done":
		Todo.CompleteTodo(id)
	default:
		fmt.Println("invalid command")
	}

}
