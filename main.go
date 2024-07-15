package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ravgus/DebtorMonitoringSystem/internal"
	"io"
	"net/http"
	"os"
)

func main() {
	defer fmt.Printf("Launch time: %v\n", internal.GetCurrentDateTime())

	internal.LoadEnv()

	internal.ValidateEnvData()

	captcha, userAgent := internal.SolveCap()
	requestBody := map[string]interface{}{
		"searchType": "1",
		"paging":     "1",
		"filter": map[string]string{
			"LastName":     os.Getenv("LAST_NAME"),
			"FirstName":    os.Getenv("FIRST_NAME"),
			"MiddleName":   os.Getenv("MIDDLE_NAME"),
			"BirthDate":    internal.DateConvert(os.Getenv("BIRTH_DATE")),
			"IdentCode":    "",
			"categoryCode": "",
			"BirthDateV":   os.Getenv("BIRTH_DATE"),
		},
		"reCaptchaToken": captcha,
	}

	jsonValue, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Cannot parse to JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://erb.minjust.gov.ua/listDebtorsEndpoint", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	response, err := internal.GetHttpClient().Do(req)
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
