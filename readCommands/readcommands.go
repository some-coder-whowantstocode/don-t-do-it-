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
		Todo.Remove(id)
	case "lt":
		arr := Todo.ListTask("./task.todo")
		Todo.ConsoleTask(arr)
	case "lc":
		arr := Todo.ListTask("./completed.todo")
		Todo.ConsoleTodo(arr)
	case "done":
		Todo.CompleteTodo(id)
	default:
		fmt.Println("invalid command")
	}

}
