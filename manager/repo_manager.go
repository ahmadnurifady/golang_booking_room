package manager

import "final-project-booking-room/repository"

type RepoManager interface {
	RoomRepo() repository.RoomRepository
}

type repoManager struct {
	infra InfraManager
}

// RoomRepo implements RepoManager.
func (r *repoManager) RoomRepo() repository.RoomRepository {
	return repository.NewRoomRepository(r.infra.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{infra: infra}
}
