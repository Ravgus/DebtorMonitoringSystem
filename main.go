package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ravgus/DebtorMonitoringSystem/internal"
	"io"
	"os"
)

func main() {
	internal.LoadEnv()

	internal.ValidateEnvData()

	captcha := internal.SolveCap()
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

	response, err := internal.GetHttpClient().Post("https://erb.minjust.gov.ua/listDebtorsEndpoint", "application/json", bytes.NewBuffer(jsonValue))
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
