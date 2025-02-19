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
	Time     string
}

func DestructureTask(task string) (string, string, string, string) {
	if task == "" {
		return "", "", "", ""
	}
	idPosition := strings.Index(task, "###ID###")
	titlePosition := strings.Index(task, "###TITLE###")
	messagePosition := strings.Index(task, "###MESSAGE###")
	timePosition := strings.Index(task, "###TIME###")
	var id, title, message, time string
	if idPosition != -1 && titlePosition != -1 {
		id = task[idPosition+8 : titlePosition]
	}
	if titlePosition != -1 && messagePosition != -1 {
		title = task[titlePosition+11 : messagePosition]
	}
	if messagePosition != -1 {
		if timePosition == -1 {
			message = task[messagePosition+13:]
		} else {
			message = task[messagePosition+13 : timePosition]
			time = task[timePosition+10:]
		}
	}

	return id, title, message, time

}

func ConsoleTask(taskarr []Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Message"})

	for _, task := range taskarr {
		table.Append([]string{task.ID, task.Title, task.Messsage})
	}

	table.Render()
}

func ConsoleTodo(taskarr []Task) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Title", "Message", "Time"})

	for _, task := range taskarr {
		table.Append([]string{task.ID, task.Title, task.Messsage, task.Time})
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

		id, title, message, t := DestructureTask(task)

		textarr = append(textarr, Task{id, title, message, t})
	}

	return textarr

}
