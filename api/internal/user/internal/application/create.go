package user

import (
	"context"
	"errors"
	"reflect"
	"time"

	"octodome.com/api/internal/user/domain"
	"octodome.com/shared/events"

	"golang.org/x/crypto/bcrypt"
)

type Create struct {
	Context  context.Context
	Name     string
	Email    string
	Password string
}

type CreateHandler struct {
	repo         domain.Repository
	eventsClient events.Client
}

func NewCreateHandler(
	repository domain.Repository,
	eventsClient events.Client,
) *CreateHandler {
	return &CreateHandler{repo: repository, eventsClient: eventsClient}
}

func (handler *CreateHandler) Handle(c Create) error {
	passwordHash, err := handler.hashPassword(c.Password)
	if err != nil {
		return err
	}

	userModel := &domain.User{
		Username:     c.Name,
		Email:        c.Email,
		PasswordHash: passwordHash,
	}

	exists, err := handler.repo.ExistsByEmailOrUsername(
		c.Context,
		userModel.Email,
		userModel.Username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email or username already exists")
	}

	userID, err := handler.repo.Create(c.Context, userModel)
	if err != nil {
		return err
	}

	handler.eventsClient.PublishEvent(
		events.EventType(reflect.TypeOf(events.UserRegistered{}).Name()),
		events.UserRegistered{
			UserID:       userID,
			Email:        userModel.Email,
			Name:         userModel.Username,
			RegisteredAt: time.Now(),
		})

	return nil
}

func (handler *CreateHandler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
