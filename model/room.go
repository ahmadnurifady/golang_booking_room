package model

import "time"

type Room struct {
	Id        string       `json:"id"`
	RoomType  string       `json:"roomType"`
	Capacity  string       `json:"capacity"`
	Facility  RoomFacility `json:"facility"`
	Status    string       `json:"status"` //untuk status hanya ada dua yaitu Available atau Booked
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
}

type RoomFacility struct {
	Id              string    `json:"id"`
	RoomId          string    `json:"roomId"`
	RoomDescription string    `json:"description"`
	Fwifi           bool      `json:"wifi"`
	Fbreakfast      bool      `json:"breakfast"`
	Fsmonking       bool      `json:"smoking"`
	Ftelevison      bool      `json:"television"`
	FcoffeMake      bool      `json:"coffe maker"`
	FbathAmenities  bool      `json:"bathroom amenities"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
