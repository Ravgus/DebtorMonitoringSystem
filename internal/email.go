package internal

import (
	"fmt"
	"github.com/go-mail/mail"
	"os"
)

func SendEmail(body string) {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER_NAME"))
	m.SetHeader("To", os.Getenv("SMTP_USER_NAME"))
	m.SetHeader("Subject", "Упс!")
	m.SetBody("text/html", body)

	port := StringToInt(os.Getenv("SMTP_PORT"))

	dialer := mail.NewDialer(
		"smtp.gmail.com",
		port,
		os.Getenv("SMTP_USER_NAME"),
		os.Getenv("SMTP_PASSWORD"),
	)
	if err := dialer.DialAndSend(m); err != nil {
		fmt.Println("Can't send email:", err)
		os.Exit(3)
	}
}

func CreateEmailBody(responseObject Response) string {
	response := "<html><body>"
	response += "<h1>" + os.Getenv("LAST_NAME") + " " + os.Getenv("FIRST_NAME") + " " + os.Getenv("MIDDLE_NAME") + "</h1>"
	response += "<p>" + "Кількість боргів: " + IntToString(responseObject.Rows)
	response += "<br>"

	for i := 0; i < len(responseObject.Result); i++ {
		response += "<br>"
		response += "<p>" + "Документ виданий: " + responseObject.Result[i].Publisher + "</p>"
		response += "<p>" + "Номер ВП: " + responseObject.Result[i].VpNum + "</p>"
		response += "<p>" + "Категорія стягнення: " + responseObject.Result[i].DeductionType + "</p>"
		response += "<p>" + "Виконавець: " + responseObject.Result[i].Executor + "</p>"
		response += "<p>" + "Пошта виконавця: " + responseObject.Result[i].ExecutorEmail + "</p>"
		response += "<p>" + "Телефон виконавця: " + responseObject.Result[i].ExecutorPhone + "</p>"
	}

	response += "<br>"
	response += "<p>" + "Для більш детальної інформації перейдіть за посиланням https://erb.minjust.gov.ua/"
	response += "</body></html>"

	return response
}
