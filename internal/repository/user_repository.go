package repository

//data layer =provide database or memory storage operations

import (
	"errors"
	"testing/internal/core"
)

type userRepository struct {
	users map[uint]*core.User
}

// NewUserRepository creates a new instance of userRepository.
func NewUserRepository() core.UserRepository {
	return &userRepository{
		users: make(map[uint]*core.User),
	}
}

func (r *userRepository) CreateUser(user *core.User) error {
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) GetUserByID(id uint) (*core.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepository) UpdateUser(user *core.User) error {
	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) DeleteUser(id uint) error {
	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}
