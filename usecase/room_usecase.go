package usecase

import (
	"final-project-booking-room/model"
	"final-project-booking-room/repository"
	"fmt"
)

type RoomUseCase interface {
	RegisterNewRoom(payload model.Room) (model.Room, error)
	FindById(id string) (model.Room, error)
	FindByRoomType(roomType string) (model.Room, error)
	ViewAllRooms() ([]model.Room, error)
	DeleteById(id string) (model.Room, error)
	UpdateById(id string, payload model.Room) (model.Room, error)
	GetRoomStatus(id string) (string, error)
	ChangeRoomStatus(id string) error
}

type roomUseCase struct {
	repo repository.RoomRepository
}

// ViewAllRooms implements RoomUseCase.
func (r *roomUseCase) ViewAllRooms() ([]model.Room, error) {
	room, err := r.repo.GetAllRoom()
	if err != nil {
		return nil, err
	}

	return room, err
}

// ChangeRoomStatus implements RoomUseCase.
func (r *roomUseCase) ChangeRoomStatus(id string) error {
	err := r.repo.ChangeStatus(id)
	if err != nil {
		panic(err)
	}

	return err
}

// GetRoomStatus implements RoomUseCase.
func (r *roomUseCase) GetRoomStatus(id string) (string, error) {
	getStatus, err := r.repo.GetStatus(id)
	if err != nil {
		return "Can't get room status", fmt.Errorf("room with id %s not found", id)
	}
	return getStatus, err
}

// FindByRoomType implements RoomUseCase.
func (r *roomUseCase) FindByRoomType(roomType string) (model.Room, error) {
	findRoom, err := r.repo.GetByRoomType(roomType)
	if err != nil {
		return model.Room{}, fmt.Errorf("room with roomType %s not found", roomType)
	}
	return findRoom, err
}

// UpdateById implements RoomUseCase.
func (r *roomUseCase) UpdateById(id string, payload model.Room) (model.Room, error) {

	return r.repo.Update(id, payload)
}

// DeleteById implements RoomUseCase.
func (r *roomUseCase) DeleteById(id string) (model.Room, error) {
	_, err := r.repo.Delete(id)
	if err != nil {
		return model.Room{}, err
	}

	return model.Room{}, err
}

// FindById implements RoomUseCase.
func (r *roomUseCase) FindById(id string) (model.Room, error) {
	findRoom, err := r.repo.Get(id)
	if err != nil {
		return model.Room{}, fmt.Errorf("room with roomType %s not found", id)
	}

	return findRoom, err
}

// RegisterNewRoom implements RoomUseCase.
func (r *roomUseCase) RegisterNewRoom(payload model.Room) (model.Room, error) {
	newRoom, err := r.repo.Create(payload)
	if err != nil {
		panic(err)
	}
	return newRoom, err
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{repo: repo}
}
