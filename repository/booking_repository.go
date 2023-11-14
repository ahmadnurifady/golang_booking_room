package repository

import (
	"database/sql"
	"final-project-booking-room/model"
	"time"
)

type BookingRepository interface {
	Create(payload model.Booking) (model.Booking, error)
	Get(id string, userId string) (model.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

// Create implements BookingRepository.
func (b *bookingRepository) Create(payload model.Booking) (model.Booking, error) {
	tx, err := b.db.Begin()
	if err != nil {
		return model.Booking{}, err
	}

	var booking model.Booking
	err = tx.QueryRow(`INSERT INTO booking (userId, updatedAt) VALUES ($1,$2) RETURNING id,userId,createdAt, updatedAt`, payload.Users.Id, time.Now()).Scan(
		&booking.Id,
		&booking.Users.Id,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return model.Booking{}, tx.Rollback()
	}

	var bookingDetails []model.BookingDetail
	for _, v := range payload.BookingDetails {
		var bookingDetail model.BookingDetail
		err = tx.QueryRow(`INSERT INTO booking_details (bookingid, roomid, bookingdate, bookingdateend, status, description, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, bookingid, roomid, bookingdate, bookingdateend, status, description, createdat, updatedat`, booking.Id, v.Rooms.Id, v.BookingDate, v.BookingDateEnd, v.Status, v.Description, time.Now()).Scan(
			&bookingDetail.Id,
			&bookingDetail.BookingId,
			&bookingDetail.Rooms.Id,
			&bookingDetail.BookingDate,
			&bookingDetail.BookingDateEnd,
			&bookingDetail.Status,
			&bookingDetail.Description,
			&bookingDetail.CreatedAt,
			&bookingDetail.UpdatedAt,
		)

		if err != nil {
			return model.Booking{}, tx.Rollback()
		}

		bookingDetail.Rooms = v.Rooms
		bookingDetails = append(bookingDetails, bookingDetail)
		if err := tx.Commit(); err != nil {
			return model.Booking{}, err
		}
	}

	booking.Users = payload.Users
	booking.BookingDetails = bookingDetails

	if err := tx.Commit(); err != nil {
		return model.Booking{}, err
	}

	return booking, err
}

// Get implements BookingRepository.
func (b *bookingRepository) Get(id string, userId string) (model.Booking, error) {
	var booking model.Booking

	err := b.db.QueryRow(`SELECT b.id, u.id, u.name, u.divisi, u.jabatan, u.email, u.role, u.createdat, u.updatedat, b.createdat, b.updatedat 
	FROM 
	booking b JOIN users u ON u.id = b.userid
	WHERE
	b.id = $1`, id).Scan(
		&booking.Id,
		&booking.Users.Id,
		&booking.Users.Name,
		&booking.Users.Divisi,
		&booking.Users.Jabatan,
		&booking.Users.Email,
		&booking.Users.Role,
		&booking.Users.CreatedAt,
		&booking.Users.UpdatedAt,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return model.Booking{}, err
	}

	var bookingDetails []model.BookingDetail
	rows, err := b.db.Query(`SELECT bd.id, bd.bookingdate, bd.bookingdateend, bd.status, bd.description, bd.createdat, bd.updatedat, r.id, r.roomtype, r.capacity, r.status, r.createdat, r.updatedat, f.id, f.roomdescription, f.fwifi, f.soundsystem, f.fprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat
	FROM 
	booking_details bd JOIN booking b ON b.id = bd.bookingid
	JOIN rooms r ON r.id = bd.roomid
	JOIN facilities f ON f.id = r.facilities 
	WHERE b.id = $1`, booking.Id)

	if err != nil {
		return model.Booking{}, err
	}

	for rows.Next() {
		var bookingDetail model.BookingDetail
		rows.Scan(
			&bookingDetail.Id,
			&bookingDetail.BookingDate,
			&bookingDetail.BookingDateEnd,
			&bookingDetail.Status,
			&bookingDetail.Description,
			&bookingDetail.CreatedAt,
			&bookingDetail.UpdatedAt,
			&bookingDetail.Rooms.Id,
			&bookingDetail.Rooms.RoomType,
			&bookingDetail.Rooms.Capacity,
			&bookingDetail.Rooms.Status,
			&bookingDetail.Rooms.CreatedAt,
			&bookingDetail.Rooms.UpdatedAt,
			&bookingDetail.Rooms.Facility.Id,
			&bookingDetail.Rooms.Facility.RoomDescription,
			&bookingDetail.Rooms.Facility.Fwifi,
			&bookingDetail.Rooms.Facility.FsoundSystem,
			&bookingDetail.Rooms.Facility.Fprojector,
			&bookingDetail.Rooms.Facility.Fchairs,
			&bookingDetail.Rooms.Facility.Ftables,
			&bookingDetail.Rooms.Facility.FsoundProof,
			&bookingDetail.Rooms.Facility.FsmonkingArea,
			&bookingDetail.Rooms.Facility.Ftelevison,
			&bookingDetail.Rooms.Facility.FAc,
			&bookingDetail.Rooms.Facility.Fbathroom,
			&bookingDetail.Rooms.Facility.FcoffeMaker,
			&bookingDetail.Rooms.Facility.UpdatedAt,
			&bookingDetail.Rooms.Facility.CreatedAt,
		)
		bookingDetails = append(bookingDetails, bookingDetail)

	}
	booking.BookingDetails = bookingDetails

	return booking, nil
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}
