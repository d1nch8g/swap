package email

import (
	"crypto/tls"
	"fmt"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	d          *gomail.Dialer
	Sender     string
	ApiAddress string
}

func New(addr, login, password, apiaddr string, port int) *Mailer {
	d := gomail.NewDialer(addr, port, login, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mes := gomail.NewMessage()

	mes.SetHeader("From", login)
	mes.SetHeader("To", login)
	mes.SetHeader("Subject", "New instance entry initialized!")
	mes.SetBody("text/html", "Hello from exchanger, new instance entry have been launched with following email!")

	err := d.DialAndSend(mes)
	if err != nil {
		echo.New().Logger.Errorf("unable to send email notification: %v", err)
	}
	return &Mailer{
		d:          d,
		Sender:     login,
		ApiAddress: apiaddr,
	}
}

func (m *Mailer) OrderCreated(email, from, to, amount, address string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Sender)
	mes.SetHeader("To", email)
	mes.SetHeader("Subject", "Order have been created.")
	mes.SetBody("text/html", fmt.Sprintf(`Created exchange order:
	Order email: %s
	Buying currency: %s
	Selling currency: %s
	Amount: %s
	Address: %s
	`, email, from, to, amount, address))

	return m.d.DialAndSend(mes)
}

func (m *Mailer) OrderFinished(email, amount, to, address string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Sender)
	mes.SetHeader("To", email)
	mes.SetHeader("Subject", "Order have been finished.")
	mes.SetBody("text/html", fmt.Sprintf(`%s %s were sent to address %s`, amount, to, address))

	return m.d.DialAndSend(mes)
}

func (m *Mailer) CancelOrder(email, from, to string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Sender)
	mes.SetHeader("To", email)
	mes.SetHeader("Subject", "Your order have been cancelled.")
	mes.SetBody("text/html", fmt.Sprintf("Your exchange order on %s to %s have been cancelled.", from, to))

	return m.d.DialAndSend(mes)
}

func (m *Mailer) UserVerifyEmail(email, uuid string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Sender)
	mes.SetHeader("To", email)
	mes.SetHeader("Subject", "Verify email address.")
	mes.SetBody("text/html", fmt.Sprintf(`Verify email address on the platform:
	follow the link:
	
	%s/api/verify/%s`, m.ApiAddress, uuid))

	return m.d.DialAndSend(mes)
}
