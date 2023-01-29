package setting

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(email *EmailSettingS) *Email {
	return &Email{
		SMTPInfo: &SMTPInfo{
			Host:     email.Host,
			Port:     email.Port,
			IsSSL:    email.IsSSL,
			UserName: email.UserName,
			Password: email.Password,
			From:     email.From,
		},
	}
}

// SendMail 发送邮件。就一个方法不放在一个单独的包里面了
func (e *Email) SendMail(to []string, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
