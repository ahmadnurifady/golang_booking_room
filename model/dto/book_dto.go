package dto

import "final-project/model"

type BookingRequestDto struct {
	Id              string                `json:"id"`
	BoookingDetails []model.BookingDetail `json:"bookingDetails" binding:"required"`
	Description     string                `json:"description"`
}
