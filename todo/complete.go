package Todo

import (
	"bufio"
	"fmt"
	"os"
)

func CompleteTodo(id string) {
	taskarr := ListTask("./task.todo")

	tempfile, err := os.OpenFile("./task.todo", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	compfile, err := os.OpenFile("./completed.todo", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println("Error while opening file")
	}

	writer := bufio.NewWriter(tempfile)
	compwriter := bufio.NewWriter(compfile)

	for i := range taskarr {
		str := FormatTask(taskarr[i].ID, taskarr[i].Title, taskarr[i].Messsage)
		if id != taskarr[i].ID {
			writer.WriteString(str)
		} else {
			compwriter.WriteString(str)
		}
	}

	writer.Flush()
	compwriter.Flush()

}
