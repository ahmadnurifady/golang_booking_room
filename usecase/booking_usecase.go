package usecase

import (
	"final-project/model"
	// "final-project/model/dto"
	"final-project/repository"
	"fmt"
)

type BookingUseCase interface {
<<<<<<< HEAD
	RegisterNewBooking(payload dto.BookingRequestDto) (model.Booking, error)
	FindById(id string) (model.Booking, error)
	ViewAllBooking() ([]model.Booking, error)
	ViewAllBookingByStatus(status string) ([]model.Booking, error)
	UpdateStatusBookAndRoom(id string, approval string) (model.Booking, error)
=======
	// RegisterNewBooking(payload dto.BookingRequestDto) (model.Booking, error)
	FindById(id string, userId string) (model.Booking, error)
>>>>>>> master
}
type bookingUseCase struct {
	repo   repository.BookingRepository
	userUC UserUseCase
	roomUC RoomUseCase
}

// UpdateStatusBookAndRoom implements BookingUseCase.
func (b *bookingUseCase) UpdateStatusBookAndRoom(id string, approval string) (model.Booking, error) {
	booking, err := b.repo.UpdateStatus(id, approval)
	if err != nil {
		return model.Booking{}, fmt.Errorf("Booking detail with ID %s not found", id)
	}
	return booking, nil
}

// ViewAllBookingByStatus implements BookingUseCase.
func (b *bookingUseCase) ViewAllBookingByStatus(status string) ([]model.Booking, error) {
	bookings, err := b.repo.GetAllByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("Failed to get data, err: %v", err)
	}
	return bookings, nil
}

// ViewAllBooking implements BookingUseCase.
func (b *bookingUseCase) ViewAllBooking() ([]model.Booking, error) {
	bookings, err := b.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("Failed to get all bookings: %v", err)
	}
	return bookings, nil
}

// FindById implements BookingUseCase.
func (b *bookingUseCase) FindById(id string) (model.Booking, error) {
	booking, err := b.repo.Get(id)
	if err != nil {
		return model.Booking{}, fmt.Errorf("Booking with ID %s not found", id)
	}
	return booking, nil
}

// RegisterNewBooking implements BookingUseCase.
<<<<<<< HEAD
func (b *bookingUseCase) RegisterNewBooking(payload dto.BookingRequestDto) (model.Booking, error) {
=======
// func (b *bookingUseCase) RegisterNewBooking(payload dto.BookingRequestDto) (model.Booking, error) {
// 	room, err := b.roomUC.FindById(payload.BoookingDetails.Room.Id)
// 	if err != nil {
// 		return model.Booking{}, err
// 	}
>>>>>>> master

// 	user, err := b.userUC.FindById(payload.UserId)
// 	if err != nil {
// 		return model.Booking{}, err
// 	}

<<<<<<< HEAD
	var bookingDetails []model.BookingDetail
	for _, v := range payload.BoookingDetails {
		room, err := b.roomUC.FindById(v.Rooms.Id)
		if err != nil {
			return model.Booking{}, err
		}

		bookingDetails = append(bookingDetails, model.BookingDetail{

			Rooms:       room,
			Description: v.Description,
			Status:      v.Status,
		})
	}

	newBookingPayload := model.Booking{
		Users:          user,
		BookingDetails: bookingDetails,
	}
=======
// 	var bookingDetails []model.BookingDetail
// 	newBookingPayload := model.Booking{
// 		Rooms: room,
// 		Users: user,
// 	}
>>>>>>> master

// 	booking, err := b.repo.Create(newBookingPayload)
// 	if err != nil {
// 		return model.Booking{}, err
// 	}

// 	return booking, nil
// }

func NewBookingUseCase(
	repo repository.BookingRepository,
	userUC UserUseCase,
	roomUC RoomUseCase,
) BookingUseCase {
	return &bookingUseCase{
		repo:   repo,
		userUC: userUC,
		roomUC: roomUC,
	}
}
