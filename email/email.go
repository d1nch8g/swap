package email

import (
	"crypto/tls"

	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"
)

type Mailer struct {
	*gomail.Dialer
}

func New(addr, login, password string) *Mailer {
	d := gomail.NewDialer(addr, 587, login, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mes := gomail.NewMessage()

	mes.SetHeader("From", login)
	mes.SetHeader("To", login)
	mes.SetAddressHeader("Cc", "support@inswap.in", "Support")
	mes.SetHeader("Subject", "New instance login initialized!")
	mes.SetBody("text/html", "Hello from exchanger!")

	err := d.DialAndSend(mes)
	if err != nil {
		echo.New().Logger.Errorf("unable to send email notification: %v", err)
	}
	return &Mailer{
		Dialer: d,
	}
}

func OrderCreated(email, from, to, amount string) {
	
}

func OrderFinished(email, from, to, amount string) {

}

func UserVerifyEmail(email, uuid string) {

}
