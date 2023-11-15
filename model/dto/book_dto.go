package dto

import "final-project-booking-room/model"

type BookingRequestDto struct {
	Id              string                `json:"id"`
	UserId          string                `json:"userId"`
	BoookingDetails []model.BookingDetail `json:"bookingDetails"`
	Description     string                `json:"description"`
}
