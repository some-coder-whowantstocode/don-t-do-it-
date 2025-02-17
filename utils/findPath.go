package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
	errorhandler "todo/errorHandler"
)

func Findpath() (string, error) {
	var path string
	lastdate, location, duration, lastdate_k_w, location_k_w, duration_k_w := "", "", "", "pre", "loc", "dur"

	f, err := os.OpenFile("./todo.config.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		errorhandler.Handler(err)
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) < 4 {
			continue
		}
		switch scanner.Text()[:3] {
		case lastdate_k_w:
			lastdate = scanner.Text()[4:]
		case location_k_w:
			location = scanner.Text()[4:]
		case duration_k_w:
			duration = scanner.Text()[4:]
		}
	}

	fmt.Println(lastdate, duration, location)

	layout := "2006-01-02"
	parsedtime, err2 := time.Parse(layout, lastdate)
	if lastdate == "" || err2 != nil {
		parsedtime = time.Now()
		lastdate = parsedtime.Format(layout)
		lastdateline := fmt.Sprintf("pre:%s\n", lastdate)
		if _, err := f.WriteString(lastdateline); err != nil {
			errorhandler.Handler(err)
			return "", err
		}
	}

	if location == "" {
		location = "todoStore"
		locationline := fmt.Sprintf("loc:%s\n", location)
		if _, err := f.WriteString(locationline); err != nil {
			errorhandler.Handler(err)
			return "", err
		}
	}

	matched, err3 := regexp.MatchString(`\d`, duration)
	if err3 != nil || !matched {
		duration = "1"
		durationline := fmt.Sprintf("dur:%s\n", duration)
		if _, err := f.WriteString(durationline); err != nil {
			errorhandler.Handler(err)
			return "", err
		}
	}

	currentTime := time.Now()
	currentDate := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	parsedDate := time.Date(parsedtime.Year(), parsedtime.Month(), parsedtime.Day(), 0, 0, 0, 0, parsedtime.Location())

	diff := int(currentDate.Sub(parsedDate).Hours() / 24)

	num, err := strconv.Atoi(duration)
	if err != nil {
		num = 1
	}

	if diff > num {
		lastdate = currentDate.Format(layout)
	}

	folderpath := fmt.Sprintf("./%s", location)

	if _, err := os.Stat(folderpath); os.IsNotExist(err) {
		if err := os.Mkdir(folderpath, os.ModePerm); err != nil {
			fmt.Println("Error creating folder:", err)
			return path, err
		}
	}

	path = fmt.Sprintf("%s/%s.txt", folderpath, lastdate)
	if _, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644); err != nil {
		fmt.Println("Error initializing file:", err)
		return path, err
	}

	return path, nil
}
