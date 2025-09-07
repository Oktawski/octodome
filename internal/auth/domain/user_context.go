package domain

type ContextKey string

const UserContextKey ContextKey = "user"

type UserContext struct {
	ID uint
}
