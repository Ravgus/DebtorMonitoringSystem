package internal

import (
	"fmt"
	"os"
	"strconv"
)

func StringToInt(data string) int {
	result, err := strconv.Atoi(data)
	if err != nil {
		fmt.Print("Can't parse string to int!")
		os.Exit(1)
	}

	return result
}

func IntToString(data int) string {
	return strconv.Itoa(data)
}
