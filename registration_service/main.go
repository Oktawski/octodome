package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"octodome.com/shared/events"
)

func main() {
	brokerURL := os.Getenv("EVENT_BROKER_URL")
	serviceURL := os.Getenv("SERVICE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	eventClient := events.NewClient(brokerURL)

	if err := eventClient.RegisterHandler(
		"registration_service",
		events.UserRegistered{}.GetEventType(),
		serviceURL+"/events",
	); err != nil {
		log.Fatalf("Failed to register handler with event broker: %v", err)
	}
	log.Printf("Registered handler for %s at %s/events", events.UserRegistered{}.GetEventType(), serviceURL)

	router := newRouter(*eventClient)

	log.Printf("Starting registration service on :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func sendEmail(p events.UserRegistered) (bool, error) {
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

	name := p.Name
	if p.Name == "" {
		name = string(p.Email)
	}

	var body strings.Builder
	body.WriteString("Hi ")
	body.WriteString(name)
	body.WriteString(",\n\nYour account has been registered successfully.\n\n")
	body.WriteString("Thanks,\nOctodome\n")

	headers := map[string]string{
		"From":         user,
		"To":           string(p.Email),
		"Subject":      "Welcome to Octodome",
		"MIME-Version": "1.0",
		"Content-Type": "text/plain; charset=UTF-8",
	}
	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(k)
		msg.WriteString(": ")
		msg.WriteString(v)
		msg.WriteString("\r\n")
	}
	msg.WriteString("\r\n")
	msg.WriteString(body.String())

	addr := fmt.Sprintf("%s:%s", host, port)

	auth := smtp.PlainAuth("", user, os.Getenv("EMAIL_SMTP_TOKEN"), host)

	if err := smtp.SendMail(addr, auth, user, []string{string(p.Email)}, []byte(msg.String())); err != nil {
		return false, err
	}
	return true, nil
}
