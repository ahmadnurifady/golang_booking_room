package repository

import (
	"database/sql"
	"errors"
	"final-project-booking-room/model"
	"final-project-booking-room/utils/common"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    UserRepository
}

func (s *UserRepositoryTestSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	assert.NoError(s.T(), err)
	s.mockDB = db
	s.mockSql = mock
	s.repo = NewUserRepository(s.mockDB)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

// func (s *UserRepositoryTestSuite) TestCreateUser_Success() {
// 	mockUser := model.User{
// 		Id:        "1",
// 		Name:      "koko",
// 		Divisi:    "HR",
// 		Jabatan:   "Senior",
// 		Email:     "koko@gmail.com",
// 		Password:  "12345",
// 		Role:      "admin",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	// Use ExpectExec for INSERT queries and check arguments
// 	// s.mockSql.ExpectExec("INSERT INTO users").WithArgs(mockUser.Id,
// 	// 	mockUser.Name, mockUser.Divisi, mockUser.Jabatan, mockUser.Email,
// 	// 	mockUser.Password, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt,
// 	// ).WillReturnResult(sqlmock.NewResult(0, 1))
// 	s.mockSql.ExpectBegin()
// 	rows := sqlmock.NewRows([]string{"id", "name", "divisi", "jabatan", "email", "password", "role", "createdat", "updatedat"}).
// 		AddRow(mockUser.Id, mockUser.Name, mockUser.Divisi, mockUser.Jabatan, mockUser.Email, mockUser.Password, mockUser.Role, mockUser.CreatedAt, mockUser.UpdatedAt)
// 	s.mockSql.ExpectQuery("INSERT INTO users").WillReturnRows(rows)

// 	s.mockSql.ExpectCommit().WillReturnError(nil)
// 	actual, err := s.repo.Create(mockUser)
// 	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet())
// 	assert.NoError(s.T(), err, "Unexpected error in Create method")
// 	assert.Equal(s.T(), mockUser.Id, actual.Id)
// }

func (s *UserRepositoryTestSuite) TestGetAllUser_Success() {

	// Define mock users
	mockUsers := []model.User{
		{
			Id:        "1",
			Name:      "John Doe",
			Divisi:    "Engineering",
			Jabatan:   "Software Engineer",
			Email:     "john.doe@example.com",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		// Add more mock users as needed
	}

	// Set expectations on the SQL mock
	s.mockSql.ExpectQuery(common.GetAllUser).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "divisi", "jabatan", "email", "role", "createdat", "updatedat"}).
			AddRow(mockUsers[0].Id, mockUsers[0].Name, mockUsers[0].Divisi, mockUsers[0].Jabatan, mockUsers[0].Email, mockUsers[0].Role, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt))

	// Call the GetAllUser method
	actualUsers, err := s.repo.GetAllUser()

	// Check for errors
	assert.NoError(s.T(), err, "Unexpected error in GetAllUser method")
	assert.Equal(s.T(), mockUsers, actualUsers, "Returned users do not match expected users")

	// Ensure all expectations were met
	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet(), "Not all SQL expectations were met")
}

func (s *UserRepositoryTestSuite) TestGetAllUser_Fail() {

	s.mockSql.ExpectQuery(common.GetAllUser).
		WillReturnError(errors.New("db error"))

	// Call the GetAllUser method
	actualUsers, err := s.repo.GetAllUser()

	expectedErrorMessage := "Expected error in GetAllUser method with 'db error' message"
	assert.Error(s.T(), err, expectedErrorMessage)

	// Check for errors

	assert.Empty(s.T(), actualUsers, "Expected empty user list for error scenario")

	// Ensure all expectations were met
	assert.Nil(s.T(), s.mockSql.ExpectationsWereMet(), "Not all SQL expectations were met")
}
