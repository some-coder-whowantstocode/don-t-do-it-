package Todo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"time"
)

func FormatCT(id, title, message string) string {
	t := time.Now().Format("2006-Jan-02")
	formattedString := fmt.Sprintf("###ID###%s###TITLE###%s###MESSAGE###%s###TIME###%s", id, title, message, t)
	encodedString := base64.StdEncoding.EncodeToString([]byte(formattedString))
	return encodedString + "\n"
}

func CompleteTodo(id string) {
	taskfile, err := os.OpenFile("./task.todo", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("Error while opening task.todo:", err)
		return
	}
	defer taskfile.Close()

	scanner := bufio.NewScanner(taskfile)

	compfile, err := os.OpenFile("./completed.todo", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error while opening completed.todo", err)
		return
	}
	defer compfile.Close()

	compwriter := bufio.NewWriter(compfile)
	defer func() {
		if err := compwriter.Flush(); err != nil {
			fmt.Println("Error while flushing compwriter:", err)
		}
	}()

	tempfile, err := os.Create("./temp.todo")
	if err != nil {
		log.Fatal("Error while creating temp.todo", err)
		return
	}
	defer tempfile.Close()

	tempwriter := bufio.NewWriter(tempfile)
	defer func() {
		if err := tempwriter.Flush(); err != nil {
			fmt.Println("Error while flushing tempwriter:", err)
		}
	}()

	found := false

	for scanner.Scan() {
		encodedString := scanner.Text()
		decodedString, err := base64.StdEncoding.DecodeString(encodedString)
		if err != nil {
			log.Fatal("Error while decoding task:", err)
		}
		Id, title, message, _ := DestructureTask(string(decodedString))
		if Id == "" || title == "" || message == "" {
			log.Fatal("task.todo file is corrupted.")
			return
		}
		if Id == id {
			found = true
			ct := FormatCT(Id, title, message)
			fmt.Println("Writing to completed.todo:", ct)
			_, err := compwriter.WriteString(ct)
			if err != nil {
				fmt.Println("Error while writing to completed.todo:", err)
			}
		} else {
			_, err := tempwriter.WriteString(encodedString + "\n")
			if err != nil {
				fmt.Println("Error while writing to temp.todo:", err)
			}
		}
	}

	if !found {
		fmt.Println("Task with ID", id, "not found.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while scanning task.todo:", err)
		return
	}

	if err := tempwriter.Flush(); err != nil {
		fmt.Println("Error while flushing tempwriter:", err)
	}

	if err := taskfile.Truncate(0); err != nil {
		log.Fatal("Error while truncating task.todo", err)
	}
	if _, err := taskfile.Seek(0, 0); err != nil {
		log.Fatal("Error while seeking task.todo", err)
	}

	tempfile.Seek(0, 0)
	tempscanner := bufio.NewScanner(tempfile)
	taskwriter := bufio.NewWriter(taskfile)
	defer func() {
		if err := taskwriter.Flush(); err != nil {
			fmt.Println("Error while flushing taskwriter:", err)
		}
	}()

	for tempscanner.Scan() {
		text := tempscanner.Text()
		fmt.Println("Writing to task.todo:", text)
		_, err := taskwriter.WriteString(text + "\n")
		if err != nil {
			fmt.Println("Error while writing to task.todo:", err)
		}
	}

	if err := tempscanner.Err(); err != nil {
		fmt.Println("Error while scanning temp.todo:", err)
		return
	}
	if err := tempfile.Truncate(0); err != nil {
		log.Fatal("Error while truncating taskfile:", err)
		return
	}
}
