package Todo

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/google/uuid"
	// "todo/utils"
)

func FormatTask(newuuid interface{}, title, msg string) string {
	formatedMessage := fmt.Sprintf("###ID###%s###TITLE###%s###MESSAGE###%s", newuuid, title, msg)
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(formatedMessage))
	return encodedMessage + "\n"
}

func AddTodo(msg, title string) {
	// path, err := utils.Findpath()

	// if err != nil {
	// 	fmt.Println("Unable to add todo")
	// 	return
	// }

	// fmt.Println(path, err)

	newuuid := uuid.New()

	f, err := os.OpenFile("./task.todo", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error while adding Todo ", err)
	}

	defer f.Close()

	if title == "" {
		title = "Task"
	}

	words := []string{"###ID###", "###TITLE###", "###MESSAGE###", "###TIME###"}
	pattern := fmt.Sprintf(`\b(%s|%s|%s|%s)\b`, words[0], words[1], words[2], words[3])

	matched, err := regexp.MatchString(pattern, title)
	matched2, err := regexp.MatchString(pattern, msg)

	if matched || matched2 {
		log.Fatal(" ###ID###,###TITLE###, ###TIME### and ###MESSAGE### are reserved keywords you can not use it.")
		return
	}

	encodedMessage := FormatTask(newuuid, title, msg)

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(encodedMessage)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	writer.Flush()
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
