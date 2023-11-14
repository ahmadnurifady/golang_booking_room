package model

import (
	"time"
)

type Booking struct {
	Id          string    `json:"bookingId"`
	UserId      User      `json:"employe"`
	RoomType    Room      `json:"RoomType"`
	CheckIn     time.Time `json:"checkIn"`
	Checkout    time.Time `json:"checkOut"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

//ketika GA acc form booking, kolom status di tabel room otomatis berubah menjadi book
