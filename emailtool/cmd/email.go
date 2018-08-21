package main

import (
	"bytes"
	"html/template"
	"log"
	"net/smtp"
	"path/filepath"
	"strings"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func (app *App) SendHtmlEmail(msg *Message) error {

	emailTemplate := filepath.Join(app.HTMLDir, "emailbody.html")
	t, err := template.ParseFiles(emailTemplate)
	if err != nil {
		return err
	}
	buffer := new(bytes.Buffer)
	//if err = t.Execute(buffer, msg); err != nil {
	if err = t.ExecuteTemplate(buffer, "emailbody", msg); err != nil {
		return err
	}

	htmlbody := buffer.String()

	body := "From: " + app.From + "\n" +
		"To: " + msg.To + "\n" +
		"Cc: " + app.Cc + "\n" +
		"Subject: " + msg.Subject + "\n" +
		MIME + "\n\n" +
		htmlbody

	//debug
	log.Println("App setting:", app)
	log.Println("body of email:", body)

	auth := smtp.PlainAuth("", app.SMTPusername, app.SMTPpassword, app.SMTPhost)
	toList := strings.Split(msg.To, ",")
	return smtp.SendMail(app.SMTPaddr, auth, app.From, toList, []byte(body))
}
