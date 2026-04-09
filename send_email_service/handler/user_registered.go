package handler

import (
	"errors"
	"html"
	"strings"

	"octodome.com/send_email_service/email"
	"octodome.com/shared/events"
)

type UserRegisteredHandler interface {
	Handle(eventID uint, p events.UserRegistered) (bool, error)
}

type userRegisteredHandler struct {
	eventsClient events.Client
}

func NewUserRegisteredHandler(eventsClient events.Client) UserRegisteredHandler {
	return &userRegisteredHandler{eventsClient: eventsClient}
}

func (handler userRegisteredHandler) Handle(eventID uint, p events.UserRegistered) (bool, error) {
	if err := handler.eventsClient.MarkEventAsProcessing(eventID); err != nil {
		return false, errors.New("failed to mark event as processing")
	}

	name := p.Name
	if name == "" {
		name = string(p.Email)
	}

	var body strings.Builder
	body.WriteString("<p>Hi ")
	body.WriteString(html.EscapeString(name))
	body.WriteString(",</p>")
	body.WriteString("<p>Your account has been registered successfully.</p>")
	body.WriteString("<p>Thanks,<br>Octodome</p>")

	if _, err := email.SendHTML("Welcome to Octodome", name, p.Email, body); err != nil {
		if err := handler.eventsClient.MarkEventAsFailed(eventID); err != nil {
			return false, errors.New("failed to mark event as failed")
		}
		return false, err
	}

	if err := handler.eventsClient.MarkEventAsProcessed(eventID); err != nil {
		return false, errors.New("failed to mark event as processed")
	}

	return true, nil
}
