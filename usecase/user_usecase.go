package usecase

import (
	"errors"
	"final-project-booking-room/model"
	"final-project-booking-room/repository"
	"final-project-booking-room/utils/common"
	"fmt"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	RegisterNewUser(payload model.User) (model.User, error)
	DeleteUser(id string) (model.User, error)
	ViewAllUser() ([]model.User, error)
	UpdateUserById(id string, payload model.User) (model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

// UpdateUserById implements UserUseCase.
func (u *userUseCase) UpdateUserById(id string, payload model.User) (model.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	var updateUser model.User

	updateUser, err = u.repo.UpdateUserById(user.Id, updateUser)
	if err != nil {
		return model.User{}, err
	}

	return updateUser, nil
}

// ViewAllUser implements UserUseCase.
func (u *userUseCase) ViewAllUser() ([]model.User, error) {
	user, err := u.repo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
}

func (u *userUseCase) DeleteUser(id string) (model.User, error) {
	_, err := u.repo.DeleteUserById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return model.User{}, nil
}

func (u *userUseCase) RegisterNewUser(payload model.User) (model.User, error) {
	if !payload.IsValidRole() {
		return model.User{}, errors.New("invalid role, role must admin or employee")
	}
	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.User{}, err
	}
	payload.Password = newPassword
	return u.repo.Create(payload)

}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
