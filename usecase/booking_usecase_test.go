package usecase

import (
	"final-project/model"
	"final-project/model/dto"
	repositorymock "final-project/unit-test/mock-test/repository-mock"
	usecasemock "final-project/unit-test/mock-test/usecase-mock"

	"testing"

	"github.com/stretchr/testify/suite"
)

type BookingUseCaseTestSuite struct {
	suite.Suite
	brm *repositorymock.BookingRepoMock
	uum *usecasemock.UserUseCaseMock
	rum *usecasemock.RoomUseCaseMock
	ues *usecasemock.EmailServiceMock
	bu  BookingUseCase
}

func (suite *BookingUseCaseTestSuite) SetupTest() {
	suite.brm = new(repositorymock.BookingRepoMock)
	suite.uum = new(usecasemock.UserUseCaseMock)
	suite.rum = new(usecasemock.RoomUseCaseMock)
	suite.ues = new(usecasemock.EmailServiceMock)
	suite.bu = NewBookingUseCase(suite.brm, suite.uum, suite.rum, suite.ues)
}

func TestBookingUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(BookingUseCaseTestSuite))
}

var mockBooking = model.Booking{
	Id: "1",
	Users: model.User{
		Id:   "1",
		Name: "Saya",
		Role: "admin",
	},
	BookingDetails: []model.BookingDetail{
		{
			Id:        "1",
			BookingId: "1",
			Rooms: model.Room{
				Id:          "1",
				RoomType:    "kamar",
				MaxCapacity: 5,
				Facility: model.RoomFacility{
					Id:              "1",
					RoomDescription: "mantap",
				},
				Status: "available",
			},
			Description: "ok",
			Status:      "pending",
		},
	},
}
var mockPayload = dto.BookingRequestDto{
	Id: "1",
	BoookingDetails: []model.BookingDetail{
		{
			Id:        "1",
			BookingId: "1",
			Rooms: model.Room{
				Id:          "5",
				RoomType:    "kolam",
				MaxCapacity: 20,
				Facility: model.RoomFacility{
					Id:              "1",
					RoomDescription: "mantap",
				},
				Status: "available",
			},
			Description: "ok",
			Status:      "pending",
		},
	},
	Description: "ok",
}

var userId = "1"

// func (suite *BookingUseCaseTestSuite) TestRegisterNewBooking() {
// 	suite.uum.On("FindById", "1").Return(mockBooking.Users, nil)

// 	var mockBookingDetails []model.BookingDetail
// 	for _, v := range mockPayload.BoookingDetails {
// 		suite.rum.On("FindById", v.Rooms.Id).Return(mockBooking.BookingDetails[0].Rooms)

// 		suite.rum.On("GetRoomStatus", v.Rooms.Id).Return(mockBooking.BookingDetails[0].Rooms.Status)

// 		mockBookingDetails = append(mockBookingDetails, mockBooking.BookingDetails[0])

// 	}

// 	mockNewBookingPayload := model.Booking{
// 		Users:          mockBooking.Users,
// 		BookingDetails: mockBookingDetails,
// 	}

// 	suite.brm.On("Create", mockNewBookingPayload).Return(mockBooking, nil)
// 	actual, err := suite.bu.RegisterNewBooking(mockPayload, userId)
// 	fmt.Println("actual:", actual)
// 	assert.Nil(suite.T(), err)
// 	assert.NoError(suite.T(), err)
// }
