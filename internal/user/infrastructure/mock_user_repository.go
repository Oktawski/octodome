package infra

import (
	"errors"
	domain "octodome/internal/user/domain"
)

type MockUserRepository struct {
	users  map[uint]*domain.User
	nextID uint
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  make(map[uint]*domain.User),
		nextID: 1,
	}
}

func (m *MockUserRepository) GetByID(id uint) (*domain.User, error) {
	user, ok := m.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) Create(user *domain.User) error {
	user.ID = m.nextID
	m.users[user.ID] = user
	m.nextID++
	return nil
}

func (m *MockUserRepository) GetUserByUsername(username string) (*domain.User, error) {
	for _, user := range m.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
