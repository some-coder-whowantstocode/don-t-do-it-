package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Findlast() {
	keyword := "pre"
	f, err := os.OpenFile("./todo.config.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var timestring string
	for scanner.Scan() {
		if scanner.Text()[0:3] == keyword {
			timestring = scanner.Text()[3:]

		}
	}

	fmt.Println(timestring)
	if timestring == "" {
		layout := "2006-01-02 15:04:05"
		date, err := time.Parse(layout, time.Now().String())
		str := fmt.Sprintf("\n%s:%s\n", keyword, date.Format(layout))
		fmt.Println(str)
		if _, err := f.Write([]byte(str)); err != nil {
			f.Close()
			log.Fatal("sdfod", err)
		}
		if err != nil {
			log.Fatal(err)
		} else {

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// return nil
}
