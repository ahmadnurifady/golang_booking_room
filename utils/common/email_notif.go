package common

import (
	"net/smtp"
	"project-final/config"
	"project-final/utils/modelutil"
)

type EmailService interface {
	SendEmail(payload modelutil.BodySender) error
}

type emailService struct {
	cfg *config.Config
}

func (e *emailService) SendEmail(payload modelutil.BodySender) error {
	message := "From: " + e.cfg.EmailFrom + "\n" +
		"To: " + payload.To[0] + "\n" +
		"Subject: " + payload.Subject + "\n\n" +
		payload.Body

	auth := smtp.PlainAuth("", e.cfg.EmailFrom, e.cfg.EmailConfig.Password, e.cfg.Server)
	err := smtp.SendMail(e.cfg.Server+":"+e.cfg.EmailConfig.Port, auth, e.cfg.EmailFrom, payload.To, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

func NewEmailService(cfg *config.Config) EmailService {
	return &emailService{cfg: cfg}
}
