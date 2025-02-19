package Todo

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
)

func Remove(id string) {
	taskfile, err := os.OpenFile("task.todo", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error while opening taskfile:", err)
		return
	}
	defer func() {
		if err := taskfile.Close(); err != nil {
			log.Fatal("Error while closing taskfile:", err)
		}
	}()

	scanner := bufio.NewScanner(taskfile)

	tempfile, err := os.OpenFile("temp.todo", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error while opening tempfile:", err)
		return
	}
	defer func() {
		if err := tempfile.Close(); err != nil {
			log.Fatal("Error while closing tempfile:", err)
		}
	}()

	writer := bufio.NewWriter(tempfile)

	for scanner.Scan() {
		encodedText := scanner.Text()
		decodedText, err := base64.StdEncoding.DecodeString(string(encodedText))
		if err != nil {
			log.Fatal("Error while decoding:", err)
			return
		}
		Id, _, _, _ := DestructureTask(string(decodedText))
		if Id != id {
			if _, err := writer.WriteString(encodedText + "\n"); err != nil {
				log.Fatal("Error while writing to tempfile:", err)
				return
			}
		}
	}

	if err := writer.Flush(); err != nil {
		log.Fatal("Error while flushing writer:", err)
		return
	}

	if err := taskfile.Truncate(0); err != nil {
		log.Fatal("Error while truncating taskfile:", err)
		return
	}

	if _, err := tempfile.Seek(0, 0); err != nil {
		log.Fatal("Error while seeking tempfile:", err)
		return
	}

	tempscanner := bufio.NewScanner(tempfile)
	taskwriter := bufio.NewWriter(taskfile)

	for tempscanner.Scan() {
		text := tempscanner.Text()
		if _, err := taskwriter.WriteString(text + "\n"); err != nil {
			log.Fatal("Error while writing to taskfile:", err)
			return
		}
	}

	if err := tempscanner.Err(); err != nil {
		log.Fatal("Error while scanning tempfile:", err)
		return
	}

	if err := taskwriter.Flush(); err != nil {
		log.Fatal("Error while flushing taskwriter:", err)
	}

	if err := tempfile.Truncate(0); err != nil {
		log.Fatal("Error while truncating taskfile:", err)
		return
	}
}
