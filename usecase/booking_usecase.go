package usecase

type BookingUseCase interface{}

type bookingUseCase struct{}

func NewBookingUseCase() BookingUseCase {
	return &bookingUseCase{}
}
