package auth

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"octodome.com/api/internal/auth/internal/dependencies"
	"octodome.com/api/internal/auth/internal/domain/repository"
	"octodome.com/shared/events"
	"octodome.com/shared/valuetype"
)

type SendMagicCodeCommand struct {
	Context context.Context
	Name    string
	Email   string
}

type SendMagicCodeHandler interface {
	Handle(cmd SendMagicCodeCommand) error
}

type sendMagicCodeHandler struct {
	repo         repository.MagicCode
	eventsClient events.Client
	userReader   repository.UserReader
}

func NewSendMagicCodeHandler(deps dependencies.Container) SendMagicCodeHandler {
	return &sendMagicCodeHandler{
		repo:         deps.MagicCodeRepository,
		userReader:   deps.UserReader,
		eventsClient: deps.EventsClient,
	}
}

func (handler *sendMagicCodeHandler) Handle(cmd SendMagicCodeCommand) error {
	userExists, _ := handler.userReader.ExistsByEmailOrUsername(cmd.Context, valuetype.Email(cmd.Email), cmd.Email)
	if !userExists {
		return errors.New("user not found")
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	if err := handler.repo.Create(cmd.Context, code, cmd.Email); err != nil {
		return err
	}

	return handler.eventsClient.PublishEvent(
		events.MagicCodeRequested{}.GetEventType(),
		events.MagicCodeRequested{Name: cmd.Name, Email: valuetype.Email(cmd.Email), Code: code})
}
