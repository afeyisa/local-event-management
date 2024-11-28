package email

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"text/template"

	"gopkg.in/gomail.v2"

	"google.golang.org/api/gmail/v1"
)



func SendEmail(userid, emailFrom, emailTo, subject, tmplPath string, data map[string]string) {
	fmt.Println("sending message" )
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		fmt.Println("location error",err)
		return
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		fmt.Println(err)
		return 
	}

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())

	var buf bytes.Buffer
	if _, err := m.WriteTo(&buf); err != nil {
		fmt.Println("2",err)

		return 
	}

	message := &gmail.Message{
		Raw: base64.URLEncoding.EncodeToString(buf.Bytes()),
	}

	GmailSeverce.Users.Messages.Send(userid, message).Do()
	
}
