package service

import (
	"gopkg.in/gomail.v2"
)

type EmailSender interface {
	SendEmail(email string, body string) error
}

type emailSender struct {
}

func NewEmailSender() EmailSender {
	return &emailSender{}
}

func (s *emailSender) SendEmail(email string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "test@hotmail.com")
	m.SetHeader("To", "huffyh00@hotmail.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello que mas!")
	m.SetBody("text/html", `
	
	<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Plantilla de correo electrónico</title>
	<style>
		/* Estilos para el encabezado */
		.header {
			background-color: #f7f7f7;
			padding: 20px;
			text-align: center;
		}
		
		/* Estilos para el logo */
		.logo {
			max-width: 100px;
		}
		
		/* Estilos para el cuerpo del mensaje */
		.body {
			padding: 20px;
		}
		
		/* Estilos para el pie de página */
		.footer {
			background-color: #f7f7f7;
			padding: 20px;
			text-align: center;
		}
	</style>
</head>
<body>
	<!-- Encabezado -->
	<div class="header">
		<img src="https://www.storicard.com/_next/static/media/icon-pay-services.089b3e6d.svg" alt="Logo de tu empresa" class="logo">
		<h1>Título del correo electrónico</h1>
	</div>
	
	<!-- Cuerpo del mensaje -->
	<div class="body">
		<p>Contenido del correo electrónico</p>
		<p>...</p>
	</div>
	
	<!-- Pie de página -->
	<div class="footer">
		<p>Este es el pie de página del correo electrónico.</p>
	</div>
</body>
</html>
	`)

	d := gomail.NewDialer("smtp.freesmtpservers.com", 25, "", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}
