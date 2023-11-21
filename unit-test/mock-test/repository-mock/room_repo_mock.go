package repositorymock

import (
	"final-project-booking-room/model"

	"github.com/stretchr/testify/mock"
)

<<<<<<< HEAD
type RoomRepoMock struct {
	mock.Mock
}

func (r *RoomRepoMock) Get(id string) (model.Room, error) {
=======
type RoomRepositoryMock struct {
	mock.Mock
}

// GetByRoomType implements repository.RoomRepository.
func (r *RoomRepositoryMock) GetByRoomType(roomType string) (model.Room, error) {
	args := r.Called(roomType)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Create(payload model.Room) (model.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Get(id string) (model.Room, error) {
>>>>>>> c2c6e8bd434977b4c92dbabb5126adc7fa4d1f50
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

<<<<<<< HEAD
func (r *RoomRepoMock) ChangeStatus(id string) error {
	args := r.Called(id)
	return args.Error(0)
=======
// func (r *RoomRepositoryMock) GetByTypeRoom(roomType string) (model.Room, error) {
// 	args := r.Called(roomType)
// 	return args.Get(0).(model.Room), args.Error(1)
// }

func (r *RoomRepositoryMock) GetAllRoom() ([]model.Room, error) {
	args := r.Called()
	return args.Get(0).([]model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Delete(id string) (model.Room, error) {
	args := r.Called(id)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) Update(id string, payload model.Room) (model.Room, error) {
	args := r.Called(id, payload)
	return args.Get(0).(model.Room), args.Error(1)
}

func (r *RoomRepositoryMock) GetStatus(id string) (string, error) {
	args := r.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (r *RoomRepositoryMock) GetStatusByBd(id string) (string, error) {
	args := r.Called(id)
	return args.Get(0).(string), args.Error(1)
}

func (r *RoomRepositoryMock) ChangeStatus(id string) error {
	args := r.Called(id)
	return args.Error(1)
}

func (r *RoomRepositoryMock) GetAllRoomByStatus(status string) ([]model.Room, error) {
	args := r.Called(status)
	return args.Get(0).([]model.Room), args.Error(1)
>>>>>>> c2c6e8bd434977b4c92dbabb5126adc7fa4d1f50
}
