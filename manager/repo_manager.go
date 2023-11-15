package manager

import "final-project-booking-room/repository"

type RepoManager interface {
	RoomRepo() repository.RoomRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infra InfraManager
}

// RoomRepo implements RepoManager.
func (r *repoManager) RoomRepo() repository.RoomRepository {
	return repository.NewRoomRepository(r.infra.Conn())
}
func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
