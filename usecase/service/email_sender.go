package service

import (
	"bytes"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/ansel1/merry"
	"github.com/wrandowR/code-challenge/config"
	"github.com/wrandowR/code-challenge/domain/model"
	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendEmail(data *model.TransactionEmail) error
}

var templateName = "email.html"

type emailSender struct {
	Template *template.Template
	Email    string
}

func NewEmailSender(email string) (EmailSender, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, merry.Wrap(err)
	}
	dir = formatDir(dir)
	sourcedir := path.Join(dir, "/templates/"+templateName)

	t, err := template.ParseFiles(sourcedir)
	if err != nil {
		return nil, merry.Wrap(err)
	}

	return &emailSender{
		Template: t,
		Email:    email,
	}, nil
}

func (s *emailSender) SendEmail(data *model.TransactionEmail) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.FromEmail(), "Code Challenge")
	m.SetHeader("To", s.Email)
	m.SetHeader("Subject", "Summary of Transactions")

	parseTemplate, err := s.parseTemplate(data)
	if err != nil {
		return merry.Wrap(err)
	}

	m.SetBody("text/html", parseTemplate)

	d := gomail.NewDialer(config.SMTPHost(), config.SMTPPort(), "", "")

	if err := d.DialAndSend(m); err != nil {
		return merry.Wrap(err)
	}

	return nil
}

func (s *emailSender) parseTemplate(data *model.TransactionEmail) (string, error) {

	buf := new(bytes.Buffer)
	if err := s.Template.Execute(buf, data); err != nil {
		return "", merry.Wrap(err)
	}

	return buf.String(), nil
}

func formatDir(dir string) string {
	basedir := strings.Split(dir, "/code-challenge")[0]
	return basedir + "/code-challenge"
}
