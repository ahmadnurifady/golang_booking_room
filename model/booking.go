package model

import (
	"time"
)

type Booking struct {
	Id               string    `json:"bookingId"`
	Users            User      `json:"employe"`
	Rooms            Room      `json:"roomType"`
	BookingDateStart time.Time `json:"checkIn"`
	BookingDateEnd   time.Time `json:"checkOut"`
	Description      string    `json:"description"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

//ketika GA acc form booking, kolom status di tabel room otomatis berubah menjadi book
