package manager

import "final-project-booking-room/usecase"

type UseCaseManager interface {
	RoomUsecase() usecase.RoomUseCase
}

type useCaseManager struct {
	repo RepoManager
}

// RoomUsecase implements UseCaseManager.
func (u *useCaseManager) RoomUsecase() usecase.RoomUseCase {
	return usecase.NewRoomUseCase(u.repo.RoomRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
