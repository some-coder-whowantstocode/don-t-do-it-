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

	matched, err := regexp.MatchString(`###end###`, title)
	matched2, err := regexp.MatchString(`###END###`, msg)

	if matched || matched2 {
		log.Fatal("###END### and ###end### are reserved keywords you can not use it.")
		return
	}

	formatedMessage := fmt.Sprintf("###START######ID###%s###TITLE###%s###MESSAGE###%s###END###", newuuid, title, msg)
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(formatedMessage))

	// if _, err := f.Write([]byte(encodedMessage)); err != nil {
	// 	f.Close()
	// 	log.Fatal(err)
	// }

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(encodedMessage + "\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	writer.Flush()
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}
