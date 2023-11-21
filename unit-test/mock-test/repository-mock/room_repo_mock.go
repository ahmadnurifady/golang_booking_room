package repositorymock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type RoomRepoMock struct {
	mock.Mock
}

func (r *RoomRepoMock) Get(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepoMock) ChangeStatus(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
