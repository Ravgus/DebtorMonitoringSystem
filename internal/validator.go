package internal

import (
	"fmt"
	"os"
)

func ValidateEnvData() {
	birthDate := os.Getenv("BIRTH_DATE")
	if !IsCorrectDate(birthDate) {
		fmt.Println("BIRTH_DATE is incorrect .env")
		return
	}

	fields := [8]string{
		"LAST_NAME",
		"FIRST_NAME",
		"MIDDLE_NAME",
		"BIRTH_DATE",
		"SMTP_USER_NAME",
		"SMTP_PASSWORD",
		"SMTP_PORT",
		"CAP_SOLVER_API_KEY",
	}

	for i := 0; i < len(fields); i++ {
		if len(os.Getenv(fields[i])) == 0 {
			fmt.Println(fields[i] + " field is empty!")
			os.Exit(4)
		}
	}
}
