package domain

type Repository interface {
	GetByID(id uint) (*User, error)
	Create(user *User) error
}
