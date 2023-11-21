package usecasemock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

type RoomUseCaseMock struct {
	mock.Mock
}

func (r *RoomUseCaseMock) FindById(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomUseCaseMock) ChangeRoomStatus(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
