package service

import (
	"bytes"
	"html/template"

	"github.com/ansel1/merry"
	"github.com/wrandowR/code-challenge/config"
	"github.com/wrandowR/code-challenge/domain/model"
	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendEmail(email string, data *model.TransactionEmail) error
}

type emailSender struct {
	TransactionEmailTemplate string
}

func NewEmailSender() EmailSender {
	return &emailSender{
		TransactionEmailTemplate: "email.html",
	}
}

func (s *emailSender) SendEmail(email string, data *model.TransactionEmail) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "huffyh00@hotmail.com")
	m.SetHeader("To", "huffyh00@hotmail.com")
	m.SetHeader("Subject", "Summary of Transactions")

	parseTemplate, err := s.parseTemplate(s.TransactionEmailTemplate, data)
	if err != nil {
		return merry.Wrap(err)
	}

	m.SetBody("text/html", parseTemplate)

	//Datos aca para el smtp de email
	d := gomail.NewDialer(config.SMTPHost(), config.SMTPPort(), "", "")

	if err := d.DialAndSend(m); err != nil {
		return merry.Wrap(err)
	}

	return nil
}

func (s *emailSender) parseTemplate(templateName string, data *model.TransactionEmail) (string, error) {
	t, err := template.ParseFiles("templates/" + s.TransactionEmailTemplate)
	if err != nil {
		return "", merry.Wrap(err)
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", merry.Wrap(err)
	}

	return buf.String(), nil
}
