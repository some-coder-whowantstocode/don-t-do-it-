package main

import (
	"flag"
	readcommands "todo/readCommands"
)

func main() {
	command := flag.String("a", `actions: add: add todo, rm: remove todo, lt: list remaining todo, lc: list completed todo`, "-a=add")
	msg := flag.String("m", "Task message", `-m="task"`)
	title := flag.String("t", "Task title", `-t="title"`)

	flag.Parse()
	readcommands.ReadCommands(*command, *msg, *title)

}
