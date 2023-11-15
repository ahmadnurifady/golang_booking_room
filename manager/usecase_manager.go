package manager

import "final-project-booking-room/usecase"

type UseCaseManager interface {
	RoomUsecase() usecase.RoomUseCase
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
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
