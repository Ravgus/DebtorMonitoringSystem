package internal

import (
	"fmt"
	"os"
	"time"
)

func DateConvert(date string) string {
	layout := "02.01.2006"
	newDate, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Can't convert date!")
		os.Exit(0)
	}

	return newDate.UTC().Format("2006-01-02T00:00:00.000Z")
}
func IsCorrectDate(date string) bool {
	layout := "02.01.2006"

	_, err := time.Parse(layout, date)

	return err == nil
}
