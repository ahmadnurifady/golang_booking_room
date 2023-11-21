package usecase

import (
	"final-project-booking-room/model"
	repositorymock "project-final/unit-test/mock-test/repository-mock"

	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomUsecaseTestSuite struct {
	suite.Suite
	rrm *repositorymock.RoomRepositoryMock
	ru  RoomUseCase
}

func (suite *RoomUsecaseTestSuite) SetupTest() {
	suite.rrm = new(repositorymock.RoomRepositoryMock)
	suite.ru = NewRoomUseCase(suite.rrm)
}

func TestRoomUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoomUsecaseTestSuite))
}

var arrayMockRoom []model.Room

var mockRoom = model.Room{
	Id:          "1",
	RoomType:    "test",
	MaxCapacity: 10,
	Facility: model.RoomFacility{
		Id:               "1",
		RoomDescription:  "ruangan test",
		Fwifi:            "ada",
		FsoundSystem:     "ada",
		Fprojector:       "ada",
		FscreenProjector: "ada",
		Fchairs:          "ada",
		Ftables:          "ada",
		FsoundProof:      "ada",
		FsmonkingArea:    "ada",
		Ftelevison:       "ada",
		FAc:              "ada",
		Fbathroom:        "ada",
		FcoffeMaker:      "ada",
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
	},
	Status:    "available",
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
}

func (suite *RoomUsecaseTestSuite) TestRegisterNewRoom_Success() {
	suite.rrm.On("Create", mockRoom).Return(mockRoom, nil)
	_, err := suite.ru.RegisterNewRoom(mockRoom)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestFindById_Success() {
	suite.rrm.On("Get", mockRoom.Id).Return(mockRoom, nil)
	_, err := suite.ru.FindById(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestFindByRoomType_Success() {
	suite.rrm.On("GetByRoomType", mockRoom.RoomType).Return(mockRoom, nil)
	_, err := suite.ru.FindByRoomType(mockRoom.RoomType)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestViewAllRooms_Success() {
	suite.rrm.On("GetAllRoom").Return(arrayMockRoom, nil)
	_, err := suite.ru.ViewAllRooms()
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestDeleteById_Success() {
	suite.rrm.On("Delete", mockRoom.Id).Return(mockRoom, nil)
	_, err := suite.ru.DeleteById(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

//update

func (suite *RoomUsecaseTestSuite) TestGetRoomStatus_Success() {
	suite.rrm.On("GetStatus", mockRoom.Id).Return(mockRoom.Status, nil)
	_, err := suite.ru.GetRoomStatus(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestGetRoomStatusByBdId_Success() {
	suite.rrm.On("GetStatusByBd", mockRoom.Id).Return(mockRoom.Id, nil)
	_, err := suite.ru.GetRoomStatusByBdId(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestChangeRoomStatus_Success() {
	suite.rrm.On("ChangeStatus", mockRoom.Id).Return(mockRoom.Id, nil)
	err := suite.ru.ChangeRoomStatus(mockRoom.Id)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUsecaseTestSuite) TestGetAllRoomByStatus() {
	suite.rrm.On("GetAllRoomByStatus", mockRoom.Status).Return(arrayMockRoom, nil)
	_, err := suite.ru.GetAllRoomByStatus(mockRoom.Status)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}
