package user

import (
	"time"

	"octodome.com/api/internal/user/domain"
	eventpublisher "octodome.com/eventbroker/eventpublisher"
	"octodome.com/shared/events"

	"golang.org/x/crypto/bcrypt"
)

type Create struct {
	Name     string
	Email    string
	Password string
}

type CreateHandler struct {
	repo           domain.Repository
	eventPublisher eventpublisher.Publisher
}

func NewCreateHandler(
	repository domain.Repository,
	eventPublisher eventpublisher.Publisher,
) *CreateHandler {
	return &CreateHandler{repo: repository, eventPublisher: eventPublisher}
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

	userID, err := handler.repo.Create(userModel)
	if err != nil {
		return err
	}

	handler.eventPublisher.Publish(events.UserRegistered{
		UserID:       userID,
		Email:        c.Email,
		Name:         c.Name,
		RegisteredAt: time.Now().UTC(),
	})

	return nil
}

func (handler *CreateHandler) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
