package usecase

import "final-project-booking-room/repository"

type RoomUseCase interface{}

type roomUseCase struct {
	repo repository.RoomRepository
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{repo: repo}
}
