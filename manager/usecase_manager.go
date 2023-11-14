package manager

import "final-project-booking-room/usecase"

type UseCaseManager interface {
	UserUseCase() usecase.UserUseCase
}

type useCaseManager struct {
	repo RepoManager
}

func (u *useCaseManager) UserUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(u.repo.UserRepo())
}

func NewUseCaseManager(repo RepoManager) UseCaseManager {
	return &useCaseManager{repo: repo}
}
