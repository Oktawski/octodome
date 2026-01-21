package user

import (
	"context"

	"octodome.com/api/internal/user/domain"
	"octodome.com/shared/events"
)

type ResetPassword struct {
	Context  context.Context
	Email    string
	Password string
}

type ResetPasswordHandler struct {
	repo         domain.Repository
	eventsClient events.Client
}

func NewResetPasswordHandler(
	repo domain.Repository,
	eventsClient events.Client,
) *ResetPasswordHandler {
	return &ResetPasswordHandler{repo: repo, eventsClient: eventsClient}
}

func (handler *ResetPasswordHandler) Handle(cmd ResetPassword) error {
	return nil
}
