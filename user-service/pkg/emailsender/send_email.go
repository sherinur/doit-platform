package emailsender

import (
	gomail "gopkg.in/mail.v2"
)

func SendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "yernarbukembay@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.gmail.com", 587, "yernarbukembay@gmail.com", "wowc mpug goas wauf")
	return d.DialAndSend(m)
}
