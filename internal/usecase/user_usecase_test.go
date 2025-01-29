package usecase_test

import (
	"testing"
	"testing/internal/core"
	"testing/internal/mocks"
	"testing/internal/usecase"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockUserUseCase := usecase.NewUserUseCase(mockRepo)

	mockUser := &core.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	mockRepo.On("CreateUser", mockUser).Return(nil)

	err := mockUserUseCase.AddUser(mockUser)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	mockUserUseCase := usecase.NewUserUseCase(mockRepo)

	mockUser := &core.User{ID: 1, Name: "John Doe", Email: "johndoe@example.com"}

	mockRepo.On("GetUserByID", uint(1)).Return(mockUser, nil)

	user, err := mockUserUseCase.GetUser(uint(1))

	assert.NoError(t, err)
	assert.Equal(t, mockUser, user)
	mockRepo.AssertExpectations(t)
}

// func TestEditUser(t *testing.T) {

// 	mockRepo := new(mocks.MockUserRepository)
// 	mockUserUseCase := usecase.NewUserUseCase(mockRepo)

// 	mockUser := &core.User{
// 		ID:    1,
// 		Name:  "John Doe Updated",
// 		Email: "johndoe_updated@example.com",
// 	}

// 	mockRepo.On("UpdateUser", mockUser).Return(nil)

// 	err := mockUserUseCase.EditUser(mockUser)

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }

// func TestDeleteUser(t *testing.T) {
// 	mockRepo := new(mocks.MockUserRepository)
// 	mockUserUseCase := usecase.NewUserUseCase(mockRepo)

// 	mockRepo.On("DeleteUser", uint(1)).Return(nil)

// 	err := mockUserUseCase.DeleteUser(uint(1))

// 	assert.NoError(t, err)
// 	mockRepo.AssertExpectations(t)
// }
