package Todo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func FormatCT(id, title, message string) string {
	t := time.Now().Format("2006-January-02")
	formattedString := fmt.Sprintf("###ID###%s###TITLE###%s###MESSAGE###%s###TIME###%s", id, title, message, t)
	encodedString := base64.StdEncoding.EncodeToString([]byte(formattedString))
	return encodedString + "\n"
}

func CompleteTodo(id string) {
	taskarr := ListTask("./task.todo")

	tempfile, err := os.OpenFile("./task.todo", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening task file:", err)
		return
	}
	defer tempfile.Close()

	writer := bufio.NewWriter(tempfile)

	var strs string

	for i := range taskarr {
		if id != taskarr[i].ID {
			str := FormatTask(taskarr[i].ID, taskarr[i].Title, taskarr[i].Messsage)
			if _, err := writer.WriteString(str); err != nil {
				fmt.Println("Error writing to task file:", err)
				return
			}
		} else {
			strs = FormatCT(taskarr[i].ID, taskarr[i].Title, taskarr[i].Messsage)

		}
	}

	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing task file writer:", err)
		return
	}

	compfile, err := os.OpenFile("./completed.todo", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening completed file:", err)
		tempfile.Close()
		return
	}
	defer compfile.Close()
	compwriter := bufio.NewWriter(compfile)

	if _, err := compwriter.WriteString(strs); err != nil {
		fmt.Println("Error writing to completed file:", err)
		return
	}

	if err := compwriter.Flush(); err != nil {
		fmt.Println("Error flushing completed file writer:", err)
		return
	}

	fmt.Println("Task updated successfully.")
}
