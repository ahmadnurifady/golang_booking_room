package usecase

import (
	"errors"
	"final-project/model"
	"final-project/repository"
	"final-project/utils/common"
	"fmt"
)

type UserUseCase interface {
	FindById(id string) (model.User, error)
	RegisterNewUser(payload model.User) (model.User, error)
	DeleteUser(id string) (model.User, error)
	ViewAllUser() ([]model.User, error)
	UpdateUserById(id string, payload model.User) (model.User, error)
	FindByEmailPassword(email string, password string) (model.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindByEmailPassword(email string, password string) (model.User, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("user with email %s not found", email)
	}

	if err := common.ComparePasswordHash(user.Password, password); err != nil {
		return model.User{}, err
	}

	user.Password = ""

	return user, nil
}

// UpdateUserById implements UserUseCase.
func (u *userUseCase) UpdateUserById(id string, payload model.User) (model.User, error) {
	if !payload.IsValidRole() {
		return model.User{}, errors.New("invalid role, role must admin or employee")
	}
	newPassword, err := common.GeneratePasswordHash(payload.Password)
	if err != nil {
		return model.User{}, err
	}
	payload.Password = newPassword
	return u.repo.UpdateUserById(id, payload)
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
