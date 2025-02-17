package Todo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func consoleTodo() {

}

func ListTodo() {
	filepath := "./task.todo"
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error while listing todo: ", err)
	}

	scanner := bufio.NewScanner(f)
	var textarr []string
	for scanner.Scan() {
		ecodedString := scanner.Bytes()
		decodedString, err := base64.StdEncoding.DecodeString(string(ecodedString))
		if err != nil {
			log.Fatal(err)
			return
		}
		textarr = append(textarr, string(decodedString))
	}
	fmt.Println(textarr)
}
