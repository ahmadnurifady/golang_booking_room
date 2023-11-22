package usecase

import (
	"final-project-booking-room/model"
	repomock "final-project-booking-room/unit-test/mock-test/repository-mock"
	usecasemock "final-project-booking-room/unit-test/mock-test/usecase-mock"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUseCaseTestSuite struct {
	suite.Suite
	uc  UserUseCase
	urm *repomock.UserRepoMock
	ues *usecasemock.EmailServiceMock
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.ues = new(usecasemock.EmailServiceMock)
	suite.urm = new(repomock.UserRepoMock)
	suite.uc = NewUserUseCase(suite.urm, suite.ues)
}

var sampleMockUser []model.User

var mockUser = model.User{
	Id:        "1",
	Name:      "test",
	Divisi:    "HR",
	Jabatan:   "Senior",
	Email:     "dika@gmail.com",
	Password:  "12345",
	Role:      "admin",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
}

func (suite *UserUseCaseTestSuite) TestRegisterNewUser_Success() {
	suite.urm.On("Create", mockUser).Return(mockUser, nil)
	_, err := suite.uc.RegisterNewUser(mockUser)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *UserUseCaseTestSuite) TestViewAllUser_Success() {
	suite.urm.On("Create", mockUser).Return(mockUser, nil)

	// Act
	result, err := suite.uc.RegisterNewUser(mockUser)

	// Assert
	suite.NoError(err)
	suite.Equal(mockUser, result)
	suite.urm.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestDeleteById_Success() {
	suite.urm.On("Delete", mockUser.Id).Return(mockUser, nil)
	_, err := suite.uc.DeleteUser(mockUser.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}
