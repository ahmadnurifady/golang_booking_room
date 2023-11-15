package usecase

import (
	"final-project-booking-room/model"
	"final-project-booking-room/repository"
)

type RoomUseCase interface {
	RegisterNewRoom(payload model.Room) (model.Room, error)
	FindById(id string) (model.Room, error)
	DeleteById(id string) error
	UpdateById(id string) error
}

type roomUseCase struct {
	repo repository.RoomRepository
}

// UpdateById implements RoomUseCase.
func (r *roomUseCase) UpdateById(id string) error {
	err := r.repo.Update(id)
	if err != nil {
		panic(err)
	}
	return err
}

// DeleteById implements RoomUseCase.
func (r *roomUseCase) DeleteById(id string) error {
	err := r.repo.Delete(id)
	if err != nil {
		panic(err)
	}

	return err
}

// FindById implements RoomUseCase.
func (r *roomUseCase) FindById(id string) (model.Room, error) {
	findRoom, err := r.repo.Get(id)
	if err != nil {
		panic(err)
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
