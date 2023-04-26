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
	m.SetBody("text/html", `<!DOCTYPE html>
	<html>
	
	<head>
		<title>Resumen de Transacciones</title>
		<style type="text/css">
			table {
				border-collapse: collapse;
				width: 100%;
			}
	
			th,
			td {
				padding: 8px;
				text-align: left;
				border-bottom: 1px solid #ddd;
			}
	
			th {
				background-color: #4CAF50;
				color: white;
			}
		</style>
	</head>
	
	<body>
		<!-- Encabezado -->
		<div class="header">
			<img src="https://www.storicard.com/_next/static/media/icon-pay-services.089b3e6d.svg" alt="Logo de tu empresa"
				class="logo">
			<h1>Resumen de Transacciones</h1>
		</div>
	
		<!-- Cuerpo del mensaje -->
		<div class="body">
			<table>
				<tr>
					<th>Mes</th>
					<th>Número de transacciones</th>
				</tr>
				<tr>
					<td>Julio</td>
					<td>2</td>
				</tr>
				<tr>
					<td>Agosto</td>
					<td>2</td>
				</tr>
			</table>
			<table>
				<tr>
					<th>Total balance is</th>
					<th>Average debit amount</th>
					<th>Average credit amount</th>
				</tr>
				<tr>
					<td>39.74</td>
					<td>-15.38</td>
					<td>35.25</td>
				</tr>
			</table>
		</div>
	
		<!-- Pie de página -->
		<div class="footer">
			<p>Gracias por su atención.</p>
		</div>
	</body>
	
	</html>`)

	d := gomail.NewDialer("smtp.freesmtpservers.com", 25, "", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return nil
}
