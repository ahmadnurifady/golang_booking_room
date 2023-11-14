package usecase

import (
	"final-project-booking-room/model"
	"final-project-booking-room/model/dto"
	"final-project-booking-room/repository"
	"fmt"
)

type BookingUseCase interface {
	RegisterNewBooking(payload dto.BookingRequestDto) (model.Booking, error)
	FindById(id string, userId string) (model.Booking, error)
}
type bookingUseCase struct {
	repo   repository.BookingRepository
	userUC UserUseCase
	roomUC RoomUseCase
}

// FindById implements BookingUseCase.
func (b *bookingUseCase) FindById(id string, userId string) (model.Booking, error) {
	bill, err := b.repo.Get(id, userId)
	if err != nil {
		return model.Booking{}, fmt.Errorf("Booking with ID %s not found", id)
	}
	return bill, nil
}

// RegisterNewBooking implements BookingUseCase.
func (b *bookingUseCase) RegisterNewBooking(payload dto.BookingRequestDto) (model.Booking, error) {
	room, err := b.roomUC.FindById(payload.BoookingDetails.Rooms.Id)
	if err != nil {
		return model.Booking{}, err
	}

	user, err := b.userUC.FindById(payload.UserId)
	if err != nil {
		return model.Booking{}, err
	}

	var bookingDetails []model.BookingDetail
	newBookingPayload := model.Booking{
		Rooms: room,
		Users: user,
	}

	booking, err := b.repo.Create(newBookingPayload)
	if err != nil {
		return model.Booking{}, err
	}

	return booking, nil
}

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
