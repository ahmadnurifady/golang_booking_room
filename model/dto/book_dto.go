package dto

type BookingRequestDto struct {
	Id          string `json:"id"`
	UserId      string `json:"userId"`
	RoomId      string `json:"roomId"`
	Description string `json:"description"`
}
