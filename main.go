package main

import (
	"DebtorMonitoringSystem/internal"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot load data from .env:", err)
		return
	}

	birthDate := os.Getenv("BIRTH_DATE")
	if !internal.IsCorrectDate(birthDate) {
		fmt.Println("BIRTH_DATE is incorrect .env:", err)
		return
	}

	requestBody := map[string]interface{}{
		"searchType": "1",
		"paging":     "1",
		"filter": map[string]string{
			"LastName":     os.Getenv("LAST_NAME"),
			"FirstName":    os.Getenv("FIRST_NAME"),
			"MiddleName":   os.Getenv("MIDDLE_NAME"),
			"BirthDate":    internal.DateConvert(birthDate),
			"IdentCode":    "",
			"categoryCode": "",
			"BirthDateV":   birthDate,
		},
	}

	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Cannot parse to JSON:", err)
		return
	}

	response, err := http.Post("https://erb.minjust.gov.ua/listDebtorsEndpoint", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Cannot reach ministry:", err)
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Cannot retrieve data from ministry:", err)
		return
	}

	var responseObject internal.Response
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		fmt.Println("Cannot parse the response:", err)
		return
	}

	if !responseObject.IsSuccess {
		fmt.Println("Ministry responded with error!")
		return
	}

	if responseObject.Rows == 0 {
		fmt.Println("You are good!")
		return
	}

	internal.SendEmail(internal.CreateEmailBody(responseObject))

	fmt.Println("Ooops! Huston, we have a problem! Corresponded data was sent by email!")
}
