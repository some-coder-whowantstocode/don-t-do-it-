package main

import (
	"flag"
	readcommands "todo/readCommands"
)

func main() {
	command := flag.String("a", "action eg:(add, rm)", "-a=add")
	msg := flag.String("m", "Task message", `-m="task"`)
	flag.Parse()
	readcommands.ReadCommands(*command, *msg)

}
