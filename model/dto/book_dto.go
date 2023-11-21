package dto

import "project-final/model"

type BookingRequestDto struct {
	Id              string                `json:"id"`
	BoookingDetails []model.BookingDetail `json:"bookingDetails"`
	Description     string                `json:"description"`
}
