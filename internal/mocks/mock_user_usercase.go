package mocks

import (
	"testing/internal/core"

	"github.com/stretchr/testify/mock"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) AddUser(user *core.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) GetUser(id uint) (*core.User, error) {
	args := m.Called(id)
	return args.Get(0).(*core.User), args.Error(1)
}

func (m *MockUserUseCase) EditUser(user *core.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserUseCase) DeleteUser(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
