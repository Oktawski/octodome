package authdom

type ContextKey string

const UserContextKey ContextKey = "user"

type UserContext struct {
	ID uint
}
