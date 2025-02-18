package main

import (
	"flag"
	readcommands "todo/readCommands"
)

func main() {
	command := flag.String("a", `actions: add: add todo, rm: remove todo, lt: list remaining todo, lc: list completed todo, done: complete task by id`, "-a=add")
	msg := flag.String("m", "Task message", `-m="task"`)
	title := flag.String("t", "Task title", `-t="title"`)
	id := flag.String("id", "Task id", `-id="f80ceec0-f89c-49e0-ba57-d1c6194acea6"`)

	flag.Parse()
	readcommands.ReadCommands(*command, *msg, *title, *id)

}
