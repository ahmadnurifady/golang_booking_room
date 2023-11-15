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
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) FindById(id string) (model.User, error) {
	user, err := u.repo.GetById(id)
	if err != nil {
		return model.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	return user, nil
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
