package usecasemock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

// DeleteUser implements usecase.UserUseCase.
func (u *UserUseCaseMock) DeleteUser(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

// FindByEmailPassword implements usecase.UserUseCase.
func (u *UserUseCaseMock) FindByEmailPassword(email string, password string) (model.User, error) {
	args := u.Called(email, password)
	return args.Get(0).(model.User), args.Error(1)
}

// RegisterNewUser implements usecase.UserUseCase.
func (u *UserUseCaseMock) RegisterNewUser(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

// UpdateUserById implements usecase.UserUseCase.
func (u *UserUseCaseMock) UpdateUserById(id string, payload model.User) (model.User, error) {
	args := u.Called(id, payload)
	return args.Get(0).(model.User), args.Error(1)
}

// ViewAllUser implements usecase.UserUseCase.
func (u *UserUseCaseMock) ViewAllUser() ([]model.User, error) {
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (u *UserUseCaseMock) FindById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
