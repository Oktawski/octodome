package userinfra

import (
	"errors"
	userdom "octodome/internal/user/domain"
)

type MockUserRepository struct {
	users  map[uint]*userdom.User
	nextID uint
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make(map[uint]*userdom.User),
		nextID: 1,
	}
}

func (m *MockUserRepository) GetByID(id uint) (*userdom.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) Create(user *userdom.User) error {
	user.ID = m.nextID
	m.users[user.ID] = user
	m.nextID++
	return nil
}

func (m *MockUserRepository) GetUserByUsername(username string) (user *userdom.User, err error) {
	for _, user := range m.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
