package usecasemock

import (
	"final-project/model"
	"final-project/model/dto"

	"github.com/stretchr/testify/mock"
)

type BillUseCaseMock struct {
	mock.Mock
}

func (b *BillUseCaseMock) RegisterNewBooking(payload dto.BookingRequestDto, userId string) (model.Booking, error) {
	args := b.Called(payload, userId)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BillUseCaseMock) FindById(id string, userId string) (model.Booking, error) {
	args := b.Called(id, userId)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BillUseCaseMock) ViewAllBooking() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BillUseCaseMock) ViewAllBookingByStatus(status string) ([]model.Booking, error) {
	args := b.Called(status)
	return args.Get(0).([]model.Booking), args.Error(1)
}

func (b *BillUseCaseMock) UpdateStatusBookAndRoom(id string, approval string) (model.Booking, error) {
	args := b.Called(id, approval)
	return args.Get(0).(model.Booking), args.Error(1)
}

func (b *BillUseCaseMock) DownloadReport() ([]model.Booking, error) {
	args := b.Called()
	return args.Get(0).([]model.Booking), args.Error(1)
}

// func (b *BillUseCaseMock) GetBookStatus(id string) (string, error) {
// 	args := b.Called(id)
// 	return args.String(0), args.Error(1)
// }

// func (b *BillUseCaseMock) GetBookingDetailsByBookingID(bookingID string) ([]model.BookingDetail, error) {
// 	args := b.Called(bookingID)
// 	return args.Get(0).([]model.BookingDetail), args.Error(1)
// }
