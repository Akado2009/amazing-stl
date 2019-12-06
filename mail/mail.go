package mail

import (
	"github.com/akado2009/amazing-stl/models"
	"gopkg.in/gomail.v2"
	"log"
)

//SendEmail func...
func SendEmail(to, attachmentPath string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", models.AppConfig.Email.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Got data!")
	m.SetBody("text/html", "Aloha amigos")
	m.Attach(attachmentPath)

	d := gomail.NewDialer(
		models.AppConfig.Email.Address,
		models.AppConfig.Email.Port,
		models.AppConfig.Email.Login,
		models.AppConfig.Email.Password,
		)
	if err := d.DialAndSend(m); err != nil {
		log.Println("error sending message: ", err)
		return err
	}
	return nil
}