package usecase

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/FaisalMashuri/emailServices/models"
	"gopkg.in/gomail.v2"
)

type EmailService struct {
	config models.EmailConfig
}

func NewEmailService(cfg models.EmailConfig) *EmailService {
	return &EmailService{
		config: cfg,
	}
}

func (s *EmailService) SendEmail(emailReciever, verificationLink string) error {
	fmt.Println(s.config)
	data := struct {
		Reciever string
		Otp      string
	}{
		Reciever: emailReciever,
		Otp:      verificationLink,
	}
	var body, err = renderTemplate("template.html", data)
	if err != nil {
		log.Println("Failed to render email template:", err)
	}
	fmt.Println("EMAIL RECIEVER : ", emailReciever)
	fmt.Println("OTP : ", verificationLink)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", s.config.ConfigSenderEmail)
	mailer.SetHeader("To", emailReciever)
	//mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", body)

	dialer := gomail.NewDialer(
		s.config.ConfigSmtpHost,
		s.config.ConfigSmtpPort,
		s.config.ConfigAuthEmail,
		s.config.ConfigAuthPassword,
	)
	fmt.Println(s.config)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
	return nil
}

func renderTemplate(templateFile string, data interface{}) (string, error) {
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return "", err
	}

	var bodyContent string
	buffer := bytes.NewBufferString(bodyContent)

	err = tmpl.Execute(buffer, data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
