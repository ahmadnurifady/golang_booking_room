package usecase

import (
	repositorymock "final-project-booking-room/unit-test/mock-test/repository-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BookingUseCaseTestSuite struct {
	suite.Suite
	brm *repositorymock.BookingRepoMock
	uum *usecasemock.UserUseCaseMock
	cum *usecasemock.RoomUseCaseMock
	bu  BookingUseCase
}

func (suite *BookingUseCaseTestSuite) SetupTest() {
	suite.brm = new(repositorymock.BookingRepoMock)
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.cum = new(usecasemock.RoomUseCaseMock)
	suite.bu = NewBookingUseCase(suite.brm, suite.uum, suite.cum)
}

func TestBookingUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BookingUseCaseTestSuite))
}
