package usecase

// business logic = implement application specific logic

import (
	"testing/internal/core"
)

type UserUseCase interface {
	AddUser(user *core.User) error
	GetUser(id uint) (*core.User, error)
	EditUser(user *core.User) error
	DeleteUser(id uint) error
}

type userUseCase struct {
	repo core.UserRepository
}

func NewUserUseCase(repo core.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) AddUser(user *core.User) error {
	return u.repo.CreateUser(user)
}

func (u *userUseCase) GetUser(id uint) (*core.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userUseCase) EditUser(user *core.User) error {
	return u.repo.UpdateUser(user)
}

func (u *userUseCase) DeleteUser(id uint) error {
	return u.repo.DeleteUser(id)
}

// func (uc *UserUseCase) AddUser(user *core.User) error {
// 	existingUser, _ := uc.repo.GetByID(user.ID)
// 	if existingUser != nil {
// 		return errors.New("user already exists")
// 	}

// 	return uc.repo.Create(user)
// }

// func (uc *UserUseCase) GetUser(id uint) (*core.User, error) {
// 	return uc.repo.GetByID(id)
// }

// func (uc *UserUseCase) EditUser(user *core.User) error {
// 	return uc.repo.Update(user)
// }

// func (uc *UserUseCase) DeleteUser(id uint) error {
// 	return uc.repo.Delete(id)
// }
