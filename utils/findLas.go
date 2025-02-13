package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
	errorhandler "todo/errorHandler"
)

func Findlast() {
	lastdate, location, duration, lastdate_k_w, location_k_w, duration_k_w := "", "", "", "pre", "loc", "dur"

	f, err := os.OpenFile("./todo.config.txt", os.O_RDWR|os.O_CREATE, 0644)
	errorhandler.Handler(err)

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) < 4 {
			continue
		}
		switch scanner.Text()[:3] {
		case lastdate_k_w:
			{
				lastdate = scanner.Text()[4:]

			}
		case location_k_w:
			{
				location = scanner.Text()[4:]

			}
		case duration_k_w:
			{
				duration = scanner.Text()[4:]
			}
		}

	}

	layout := "2006-01-02"
	_, err2 := time.Parse(layout, lastdate)
	if lastdate == "" || err2 != nil {
		lastdate = time.Now().Format(layout)
	}

	if location == "" {
		location = "todoStore"
	}

	matched, err3 := regexp.Match(`/\d/g`, []byte(`duration`))

	errorhandler.Handler(err3)

	if duration == "" || !matched {
		duration = "1"
	}

	fmt.Println(lastdate, location, duration)
	path := fmt.Sprintf("./%s/%s.txt", location, lastdate)
	fmt.Println(path)
	_, err9 := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	errorhandler.Handler(err9)

	// parsedtime, err := time.Parse(layout, lastdate)
	// fmt.Println(parsedtime.Year())
}
