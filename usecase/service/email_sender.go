package repository

type EmailSender interface {
	SendEmail(email string, body string) error
}

type emailSender struct {
}

func NewEmailSender() EmailSender {
	return &emailSender{}
}

func (s *emailSender) SendEmail(email string, body string) error {
	return nil
}
