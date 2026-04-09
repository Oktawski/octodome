package auth

import (
	"context"
	"errors"

	"octodome.com/api/internal/auth/internal/dependencies"
	"octodome.com/api/internal/auth/internal/domain"
	domainrepo "octodome.com/api/internal/auth/internal/domain/repository"
)

type AuthenticateMagicCodeCommand struct {
	Context context.Context
	Email   string
	Code    string
}

type AuthenticateMagicCodeHandler interface {
	Handle(cmd AuthenticateMagicCodeCommand) (string, error)
}

type authenticateMagicCodeHandler struct {
	repo           domainrepo.MagicCode
	userReader     domainrepo.UserReader
	tokenGenerator domain.AuthTokenGenerator
}

func NewAuthenticateMagicCodeHandler(deps dependencies.Container) AuthenticateMagicCodeHandler {
	return &authenticateMagicCodeHandler{
		repo:           deps.MagicCodeRepository,
		userReader:     deps.UserReader,
		tokenGenerator: deps.TokenGenerator,
	}
}

func (handler *authenticateMagicCodeHandler) Handle(cmd AuthenticateMagicCodeCommand) (string, error) {
	userAuthDTO, err := handler.userReader.GetUserAuthDTO(cmd.Context, cmd.Email)
	if err != nil {
		return "", err
	}

	code, err := handler.repo.GetByEmailAndCode(cmd.Context, cmd.Email, cmd.Code)
	if err != nil {
		return "", err
	}

	if code != cmd.Code {
		return "", errors.New("invalid code")
	}

	if err := handler.repo.DeleteByEmail(cmd.Context, cmd.Email); err != nil {
		return "", err
	}

	return handler.tokenGenerator.GenerateToken(userAuthDTO)
}
