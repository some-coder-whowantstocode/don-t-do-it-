package Todo

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	ID       string
	Title    string
	Messsage string
}

func DestructureTask(task string) (string, string, string) {
	trimmedString := task[11:]
	trimmedString = trimmedString[:len(trimmedString)-9]
	idPosition := strings.Index(trimmedString, "###ID###")
	titlePosition := strings.Index(trimmedString, "###TITLE###")
	messagePosition := strings.Index(trimmedString, "###MESSAGE###")
	id := trimmedString[idPosition+8 : titlePosition]
	title := trimmedString[titlePosition+11 : messagePosition]
	message := trimmedString[messagePosition+13:]

	return id, title, message

}

func ConsoleTodo(taskarr []Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Message"})

	for _, task := range taskarr {
		table.Append([]string{task.ID, task.Title, task.Messsage})
	}

	table.Render()
}

func ListTask(filepath string) []Task {
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error while listing todo: ", err)
	}

	scanner := bufio.NewScanner(f)
	var textarr []Task
	for scanner.Scan() {
		ecodedString := scanner.Bytes()
		decodedString, err := base64.StdEncoding.DecodeString(string(ecodedString))
		if err != nil {
			log.Fatal(err)
			return nil
		}

		task := string(decodedString)

		if task == "" {
			continue
		}

		id, title, message := DestructureTask(task)

		textarr = append(textarr, Task{id, title, message})
	}

	return textarr

}
