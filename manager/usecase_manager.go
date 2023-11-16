package manager

import (
	"final-project-booking-room/usecase"
)

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	RoomUsecase() usecase.RoomUseCase
	BookingUsecase() usecase.BookingUseCase
}

type useCaseManager struct {
	repo RepoManager
}

// BookingUsecase implements UseCaseManager.
func (u *useCaseManager) BookingUsecase() usecase.BookingUseCase {
	return usecase.NewBookingUseCase(u.repo.BookingRepo(), u.UserUseCase(), u.RoomUsecase())
}

// RoomUsecase implements UseCaseManager.
func (u *useCaseManager) RoomUsecase() usecase.RoomUseCase {
	return usecase.NewRoomUseCase(u.repo.RoomRepo())
}
func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
