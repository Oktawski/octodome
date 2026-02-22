package settings

import (
	"context"

	authdom "octodome.com/api/internal/auth/domain"
	"octodome.com/api/internal/settings/internal/domain"
)

type Upsert struct {
	Context     context.Context
	UserContext authdom.UserContext
	Name        string
	Value       string
}

type UpsertHandler struct {
	repo domain.Repository
}

func NewUpsertHandler(repo domain.Repository) *UpsertHandler {
	return &UpsertHandler{repo: repo}
}

func (h *UpsertHandler) Handle(c Upsert) error {
	setting := domain.NewSetting(c.Name, c.Value)
	return h.repo.Set(c.Context, *setting)
}
