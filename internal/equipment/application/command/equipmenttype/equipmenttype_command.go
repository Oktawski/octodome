package eqtypecmd

import authdom "octodome/internal/auth/domain"

type Create struct {
	Name        string
	UserContext authdom.UserContext
}

type Update struct {
	ID          uint
	Name        string
	UserContext authdom.UserContext
}

type Delete struct {
	ID          uint
	UserContext authdom.UserContext
}
