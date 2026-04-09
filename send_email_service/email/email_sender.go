package email

import (
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"octodome.com/shared/valuetype"
)

func SendHTML(
	subject string,
	toName string,
	toEmail valuetype.Email,
	emailBody strings.Builder,
) (bool, error) {
	host := os.Getenv("EMAIL_SMTP_HOST")
	if host == "" {
		return false, errors.New("EMAIL_SMTP_HOST is required")
	}

	port := os.Getenv("EMAIL_SMTP_PORT")
	if port == "" {
		return false, errors.New("EMAIL_SMTP_PORT is required")
	}

	user := os.Getenv("EMAIL_SMTP_USER")
	if user == "" {
		return false, errors.New("EMAIL_SMTP_USER is required")
	}

	headers := map[string]string{
		"From":         user,
		"To":           string(toEmail),
		"Subject":      subject,
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(k)
		msg.WriteString(": ")
		msg.WriteString(v)
		msg.WriteString("\r\n")
	}
	msg.WriteString("\r\n")
	msg.WriteString(emailBody.String())

	addr := fmt.Sprintf("%s:%s", host, port)

	auth := smtp.PlainAuth("", user, os.Getenv("EMAIL_SMTP_TOKEN"), host)

	if err := smtp.SendMail(addr, auth, user, []string{string(toEmail)}, []byte(msg.String())); err != nil {
		return false, err
	}
	return true, nil
}
