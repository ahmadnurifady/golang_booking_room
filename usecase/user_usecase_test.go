package usecase

import (
	repomock "final-project-booking-room/mock/repository_mock"
	usecasemock "final-project-booking-room/mock/usecase_mock"

	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	uc  UserUseCase
	urm *repomock.UserRepoMock
	ucm *usecasemock.UserUseCaseMock
	ues *usecasemock.MockEmailService
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.ucm = new(usecasemock.UserUseCaseMock)
	suite.ues = new(usecasemock.MockEmailService)
	suite.urm = new(repomock.UserRepoMock)
	// suite.uc = NewUserUseCase(suite.urm, suite.ues)
}
