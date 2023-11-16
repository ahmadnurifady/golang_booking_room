package usecase

import (
	"final-project/model"
	"final-project/repository"
	"fmt"
)

type RoomUseCase interface {
	RegisterNewRoom(payload model.Room) (model.Room, error)
	FindById(id string) (model.Room, error)
	FindByRoomType(roomType string) (model.Room, error)
	ViewAllRooms() ([]model.Room, error)
	DeleteById(id string) (model.Room, error)
	UpdateById(id string, payload model.Room) (model.Room, error)
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
