package userdom

type UserRepository interface {
	GetByID(id uint) (*User, error)
	Create(user *User) error
}
