package repositorymock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (u *UserRepoMock) GetById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) Create(payload model.User) (model.User, error) {
	args := u.Called(payload)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepoMock) UpdateUserById(userId string, payload model.User) (model.User, error) {
	args := u.Called(userId, payload)
	return args.Get(0).(model.User), args.Error(1)
}
func (u *UserRepoMock) DeleteUserById(id string) (model.User, error) {
	args := u.Called(id)
	return args.Get(0).(model.User), args.Error(1)
}
func (u *UserRepoMock) GetAllUser() ([]model.User, error) {
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}
func (u *UserRepoMock) GetByEmail(email string) (model.User, error) {
	args := u.Called(email)
	return args.Get(0).(model.User), args.Error(1)
}
