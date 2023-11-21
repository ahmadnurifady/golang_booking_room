package repository

import (
	"database/sql"
	"project-final/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RoomRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    RoomRepository
}

func (suite *RoomRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(suite.T(), err)
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewRoomRepository(suite.mockDB)
}

func TestRoomRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomRepositoryTestSuite))
}

func (suite *RoomRepositoryTestSuite) TestCreateRoom_Success() {
	mockRoom := model.Room{
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

	rows := sqlmock.NewRows([]string{"id", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "createdat", "updatedat"}).AddRow(
		mockRoom.Facility.Id,
		mockRoom.Facility.RoomDescription,
		mockRoom.Facility.Fwifi,
		mockRoom.Facility.FsoundSystem,
		mockRoom.Facility.Fprojector,
		mockRoom.Facility.FscreenProjector,
		mockRoom.Facility.Fchairs,
		mockRoom.Facility.Ftables,
		mockRoom.Facility.FsoundProof,
		mockRoom.Facility.FsmonkingArea,
		mockRoom.Facility.Ftelevison,
		mockRoom.Facility.FAc,
		mockRoom.Facility.Fbathroom,
		mockRoom.Facility.FcoffeMaker,
		mockRoom.Facility.CreatedAt,
		mockRoom.Facility.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO facilities").WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"id", "roomtype", "capacity", "status", "createdat", "updatedat"}).AddRow(
		mockRoom.Id,
		mockRoom.RoomType,
		mockRoom.MaxCapacity,
		mockRoom.Status,
		mockRoom.CreatedAt,
		mockRoom.UpdatedAt)
	suite.mockSql.ExpectQuery("INSERT INTO rooms").WillReturnRows(rows)

	suite.mockSql.ExpectCommit()
	actual, err := suite.repo.Create(mockRoom)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom.Id, actual.Id)

}

func (suite *RoomRepositoryTestSuite) TestGetByRoomType() {
	mockRoom := model.Room{
		Id:          "1",
		RoomType:    "room test",
		MaxCapacity: 10,
		Facility: model.RoomFacility{
			Id:               "101",
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

	rows := sqlmock.NewRows([]string{
		"id", "roomtype", "capacity",
		"facilityid", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "updatedat", "createdat",
		"status", "createdat", "updatedat"}).AddRow(
		mockRoom.Id,
		mockRoom.RoomType,
		mockRoom.MaxCapacity,
		mockRoom.Facility.Id,
		mockRoom.Facility.RoomDescription,
		mockRoom.Facility.Fwifi,
		mockRoom.Facility.FsoundSystem,
		mockRoom.Facility.Fprojector,
		mockRoom.Facility.FscreenProjector,
		mockRoom.Facility.Fchairs,
		mockRoom.Facility.Ftables,
		mockRoom.Facility.FsoundProof,
		mockRoom.Facility.FsmonkingArea,
		mockRoom.Facility.Ftelevison,
		mockRoom.Facility.FAc,
		mockRoom.Facility.Fbathroom,
		mockRoom.Facility.FcoffeMaker,
		mockRoom.Facility.UpdatedAt,
		mockRoom.Facility.CreatedAt,
		mockRoom.Status,
		mockRoom.CreatedAt,
		mockRoom.UpdatedAt,
	)

	suite.mockSql.ExpectQuery(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities WHERE r.roomtype = \$1;`).
		WithArgs("room test").
		WillReturnRows(rows)

	result, err := suite.repo.GetByRoomType("room test")

	// assert.NoError(suite.T(), err)
	// assert.Equal(suite.T(), mockRoom, result)
	// assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())

	assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockRoom, result)

	// suite.mockSql.ExpectQuery("^SELECT .*").WithArgs("roomtype_value").WillReturnRows(sqlmock.NewRows(
	// 	[]string{"id", "roomtype", "capacity", "facilitie.id", "facilitie.roomdescription", "facilitie.fwifi", "facilitie.fsoundsystem", "facilitie.fprojector", "facilitie.fscreenprojector", "facilitie.fchairs", "facilitie.ftables", "facilitie.fsoundproof", "facilitie.fsmonkingarea", "facilitie.ftelevison", "facilitie.fac", "facilitie.fbathroom", "facilitie.fcoffemaker", "facilitie.createdat", "facilitie.updatedat", "status", "createdat", "updatedat"}).AddRow(mockRoom.Id, mockRoom.RoomType, mockRoom.MaxCapacity,
	// 	mockRoom.Facility.Id, mockRoom.Facility.RoomDescription, mockRoom.Facility.Fwifi,
	// 	mockRoom.Facility.FsoundSystem, mockRoom.Facility.Fprojector, mockRoom.Facility.FscreenProjector,
	// 	mockRoom.Facility.Fchairs, mockRoom.Facility.Ftables, mockRoom.Facility.FsoundProof,
	// 	mockRoom.Facility.FsmonkingArea, mockRoom.Facility.Ftelevison, mockRoom.Facility.FAc,
	// 	mockRoom.Facility.Fbathroom, mockRoom.Facility.FcoffeMaker, mockRoom.Facility.CreatedAt,
	// 	mockRoom.Facility.UpdatedAt, mockRoom.Status, mockRoom.CreatedAt, mockRoom.UpdatedAt))

	// actual, err := suite.repo.GetByRoomType(mockRoom.RoomType)

	// assert.NoError(suite.T(), suite.mockSql.ExpectationsWereMet())
	// assert.Nil(suite.T(), err)
	// assert.NoError(suite.T(), err)
	// assert.Len(suite.T(), actual, 1)

}

func (suite *RoomRepositoryTestSuite) TestGetRoom_Success() {

	// expectedRoom := model.Room{
	// 	Id:          "1",
	// 	RoomType:    "test",
	// 	MaxCapacity: 10,
	// 	Facility: model.RoomFacility{
	// 		Id:               "1",
	// 		RoomDescription:  "ruangan test",
	// 		Fwifi:            "",
	// 		FsoundSystem:     "",
	// 		Fprojector:       "",
	// 		FscreenProjector: "",
	// 		Fchairs:          "",
	// 		Ftables:          "",
	// 		FsoundProof:      "",
	// 		FsmonkingArea:    "",
	// 		Ftelevison:       "",
	// 		FAc:              "",
	// 		Fbathroom:        "",
	// 		FcoffeMaker:      "",
	// 		CreatedAt:        time.Time{},
	// 		UpdatedAt:        time.Time{},
	// 	},
	// 	Status:    "available",
	// 	CreatedAt: time.Time{},
	// 	UpdatedAt: time.Time{},
	// }

	// rows := sqlmock.NewRows([]string{"id", "roomtype", "capacity", "id", "roomdescription", "fwifi", "fsoundsystem", "fprojector", "fscreenprojector", "fchairs", "ftables", "fsoundproof", "fsmonkingarea", "ftelevison", "fac", "fbathroom", "fcoffemaker", "createdat", "updatedat", "status", "createdat", "updatedat"}).
	// 	AddRow(expectedRoom.Id, expectedRoom.RoomType, expectedRoom.MaxCapacity,
	// 		expectedRoom.Facility.Id, expectedRoom.Facility.RoomDescription, expectedRoom.Facility.Fwifi,
	// 		expectedRoom.Facility.FsoundSystem, expectedRoom.Facility.Fprojector, expectedRoom.Facility.FscreenProjector,
	// 		expectedRoom.Facility.Fchairs, expectedRoom.Facility.Ftables, expectedRoom.Facility.FsoundProof,
	// 		expectedRoom.Facility.FsmonkingArea, expectedRoom.Facility.Ftelevison, expectedRoom.Facility.FAc,
	// 		expectedRoom.Facility.Fbathroom, expectedRoom.Facility.FcoffeMaker, expectedRoom.Facility.CreatedAt,
	// 		expectedRoom.Facility.UpdatedAt, expectedRoom.Status, expectedRoom.CreatedAt, expectedRoom.UpdatedAt)

}
