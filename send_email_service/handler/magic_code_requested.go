package handler

import (
	"errors"
	"html"
	"strings"

	"octodome.com/send_email_service/email"
	"octodome.com/shared/events"
)

type MagicCodeRequestedHandler interface {
	Handle(eventID uint, p events.MagicCodeRequested) (bool, error)
}

type magicCodeRequestedHandler struct {
	eventsClient events.Client
}

func NewMagicCodeRequestedHandler(eventsClient events.Client) MagicCodeRequestedHandler {
	return &magicCodeRequestedHandler{eventsClient: eventsClient}
}

func (handler magicCodeRequestedHandler) Handle(eventID uint, p events.MagicCodeRequested) (bool, error) {
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
	body.WriteString("<p>Your code to sign in is:</p>")
	body.WriteString("<p><strong>")
	body.WriteString(html.EscapeString(p.Code))
	body.WriteString("</strong></p>")
	body.WriteString("<p>Thanks,<br>Octodome</p>")

	if _, err := email.SendHTML("Single Sign On - Magic Code", name, p.Email, body); err != nil {
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
