package email

import (
	"crypto/tls"
	"fmt"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	*gomail.Dialer
	Name string
}

func New(addr, login, password, name string) *Mailer {
	d := gomail.NewDialer(addr, 587, login, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mes := gomail.NewMessage()

	mes.SetHeader("From", login)
	mes.SetHeader("To", login)
	mes.SetHeader("Subject", "New instance login initialized!")
	mes.SetBody("text/html", "Hello from exchanger!")

	err := d.DialAndSend(mes)
	if err != nil {
		echo.New().Logger.Errorf("unable to send email notification: %v", err)
	}
	return &Mailer{
		Dialer: d,
		Name:   name,
	}
}

func (m *Mailer) OrderCreated(email, from, to, amount, address string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Name)
	mes.SetHeader("To", email)
	mes.SetHeader("Subject", "Order have been created.")
	mes.SetBody("text/html", fmt.Sprintf(`Created exchange order:
	Order email: %s
	Buying currency: %s
	Selling currency: %s
	Amount: %s
	Address: %s
	`, email, from, to, amount))

	return m.DialAndSend(mes)
}

func (m *Mailer) OrderFinished(email, amount, to, address string) error {
	mes := gomail.NewMessage()

	mes.SetHeader("From", m.Name)
	mes.SetHeader("To", email)
	mes.SetHeader("Subject", "Order have been finished.")
	mes.SetBody("text/html", fmt.Sprintf(`%s %s were sent to address %s`, amount, to, address))

	return m.DialAndSend(mes)
}

func (m *Mailer) UserVerifyEmail(email, uuid string) {

}
