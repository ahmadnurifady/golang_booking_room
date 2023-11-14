package repository

type BookingRepository interface{}

type bookingRepository struct{}

func NewBookingRepository() BookingRepository {
	return &bookingRepository{}
}
