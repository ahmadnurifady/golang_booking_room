package manager

import (
	"project-final/usecase"
	"project-final/utils/common"
)

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
	RoomUsecase() usecase.RoomUseCase
	BookingUsecase() usecase.BookingUseCase
}

type useCaseManager struct {
	repo  RepoManager
	email common.EmailService
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
	return usecase.NewUserUseCase(u.repo.UserRepo(), u.email)
}

func NewUseCaseManager(repo RepoManager, email common.EmailService) UseCaseManager {
	return &useCaseManager{repo: repo, email: email}
}
