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
	err = tx.QueryRow(`INSERT INTO booking (userId, roomId, bookingdatestart, bookingdateend, description, status, updatedAt) VALUES ($1,$2,$3,$4, $5, $6, $7) RETURNING id,userId, roomId, bookingdatestart, bookingdateend, description, status, createdAt, updatedAt`, payload.Users.Id, payload.Rooms.Id, payload.BookingDateStart, payload.BookingDateEnd, payload.Description, payload.Status, time.Now()).Scan(
		&booking.Id,
		&booking.Users.Id,
		&booking.Rooms.Id,
		&booking.BookingDateStart,
		&booking.BookingDateEnd,
		&booking.Description,
		&booking.Status,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return model.Booking{}, tx.Rollback()
	}

	booking.Rooms = payload.Rooms
	booking.Users = payload.Users

	return booking, err
}

// Get implements BookingRepository.
func (b *bookingRepository) Get(id string, userId string) (model.Booking, error) {
	var booking model.Booking
	//query r.facilitiesid cek, takut nya beda
	err := b.db.QueryRow(`
		SELECT b.id, u.id, u.name, u.divisi, u.jabatan, u.email, u.role, u.createdAt, u.updatedat,
		r.id , r.roomtype, r.capacity, 
		r.facilitiesid, 
		f.roomdescription, f.fwifi, f.fsounsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.soundproof, f.fsmonkingarea, f.ftelevision, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat,
		r.status, r.createdat, r.updatedat, b.bookingdatestart, b.bookingdateend, b.description, b.status, b.createdat, b.updatedat FROM booking b JOIN users u ON u.id = b.userid JOIN rooms r ON r.id = b.roomid WHERE b.id = $1 OR b.user_id = $2`, id, userId).
		Scan(
			&booking.Id,
			&booking.Users.Id,
			&booking.Users.Name,
			&booking.Users.Divisi,
			&booking.Users.Jabatan,
			&booking.Users.Email,
			&booking.Users.Role,
			&booking.Users.CreatedAt,
			&booking.Users.UpdatedAt,
			&booking.Rooms.Id,
			&booking.Rooms.RoomType,
			&booking.Rooms.Capacity,

			&booking.Rooms.Facility.Id,
			&booking.Rooms.Facility.RoomDescription,
			&booking.Rooms.Facility.Fwifi,
			&booking.Rooms.Facility.FsoundSystem,
			&booking.Rooms.Facility.Fprojector,
			&booking.Rooms.Facility.FscreenProjector,
			&booking.Rooms.Facility.Fchairs,
			&booking.Rooms.Facility.Ftables,
			&booking.Rooms.Facility.FsoundProof,
			&booking.Rooms.Facility.FsmonkingArea,
			&booking.Rooms.Facility.Ftelevison,
			&booking.Rooms.Facility.FAc,
			&booking.Rooms.Facility.Fbathroom,
			&booking.Rooms.Facility.FcoffeMaker,
			&booking.Rooms.Facility.CreatedAt,
			&booking.Rooms.Facility.UpdatedAt,

			&booking.Rooms.Status,
			&booking.Rooms.CreatedAt,
			&booking.Rooms.UpdatedAt,
			&booking.BookingDateStart,
			&booking.BookingDateEnd,
			&booking.Description,
			&booking.Status,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)

	if err != nil {
		return model.Booking{}, err
	}

	return booking, nil

}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db: db}
}
