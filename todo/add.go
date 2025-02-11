package Todo

import (
	"log"
	"os"
)

func AddTodo(msg string) {
	f, err := os.OpenFile("./todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error while adding Todo", err)
	}

	if _, err := f.Write([]byte(msg)); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
