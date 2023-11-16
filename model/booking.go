package model

import (
	"time"
)

type Booking struct {
	Id             string          `json:"bookingId"`
	Users          User            `json:"employe"`
	BookingDetails []BookingDetail `json:"bookingDetails"`
	CreatedAt      time.Time       `json:"createdAt"`
	UpdatedAt      time.Time       `json:"updatedAt"`
}

type BookingDetail struct {
<<<<<<< HEAD
	Id             string    `json:"id"`
	BookingId      string    `json:"bookingId"`
	Rooms          Room      `json:"rooms"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	BookingDate    time.Time `json:"bookingDate"`
	BookingDateEnd time.Time `json:"bookingDateEnd"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
=======
	Id               string    `json:"id"`
	BookingId        string    `json:"bookingId"`
	RoomType         Room      `json:"RoomType"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	BookingDateStart time.Time `json:"checkIn"`
	BookingDateEnd   time.Time `json:"checkOut"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
>>>>>>> master
}
