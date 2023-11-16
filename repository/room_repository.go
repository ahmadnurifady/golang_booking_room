package repository

import (
	"database/sql"
	"final-project/model"
	"time"
)

type RoomRepository interface {
	Create(payload model.Room) (model.Room, error)
	Get(id string) (model.Room, error)
	GetByRoomType(roomType string) (model.Room, error)
	GetAllRoom() ([]model.Room, error)
	Delete(id string) error
	Update(id string) (model.Room, error)
}

type roomRepository struct {
	db *sql.DB
}

// GetAllRoom implements RoomRepository.
func (r *roomRepository) GetAllRoom() ([]model.Room, error) {
	var rooms []model.Room

	rows, err := r.db.Query(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var room model.Room
		err := rows.Scan(
			&room.Id,
			&room.RoomType,
			&room.MaxCapacity,
			&room.Facility.Id,
			&room.Facility.RoomDescription,
			&room.Facility.Fwifi,
			&room.Facility.FsoundSystem,
			&room.Facility.Fprojector,
			&room.Facility.FscreenProjector,
			&room.Facility.Fchairs,
			&room.Facility.Ftables,
			&room.Facility.FsoundProof,
			&room.Facility.FsmonkingArea,
			&room.Facility.Ftelevison,
			&room.Facility.FAc,
			&room.Facility.Fbathroom,
			&room.Facility.FcoffeMaker,
			&room.Facility.CreatedAt,
			&room.Facility.UpdatedAt,
			&room.Status,
			&room.CreatedAt,
			&room.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	if err != nil {
		return nil, err
	}

	return rooms, err
}

// GetByRoomType implements RoomRepository.
func (r *roomRepository) GetByRoomType(roomType string) (model.Room, error) {
	var room model.Room
	err := r.db.QueryRow(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities WHERE r.roomtype = $1;`, roomType).Scan(
		&room.Id,
		&room.RoomType,
		&room.MaxCapacity,
		&room.Facility.Id,
		&room.Facility.RoomDescription,
		&room.Facility.Fwifi,
		&room.Facility.FsoundSystem,
		&room.Facility.Fprojector,
		&room.Facility.FscreenProjector,
		&room.Facility.Fchairs,
		&room.Facility.Ftables,
		&room.Facility.FsoundProof,
		&room.Facility.FsmonkingArea,
		&room.Facility.Ftelevison,
		&room.Facility.FAc,
		&room.Facility.Fbathroom,
		&room.Facility.FcoffeMaker,
		&room.Facility.CreatedAt,
		&room.Facility.UpdatedAt,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return model.Room{}, err
	}
	return room, err
}

// Update implements RoomRepository.
func (r *roomRepository) Update(id string) (model.Room, error) {
	var room model.Room
	var facilitie model.RoomFacility

	err := r.db.QueryRow(`UPDATE rooms SET roomtype = $1, capacity = $2, status = $3, updatedat = $4 WHERE id = $5 RETURNING facilities, id, roomtype, capacity, status, createdat, updatedat`, room.RoomType, room.MaxCapacity, room.Status, time.Now(), id).Scan(
		&room.Facility.Id,
		&room.Id,
		&room.RoomType,
		&room.MaxCapacity,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	facilitie.Id = room.Facility.Id

	err = r.db.QueryRow(`UPDATE facilities SET roomdescription = $1, fwifi = $2, fsoundsystem = $3, fprojector = $4, fscreenprojector = $5, fchairs = $6, ftables = $7, fsoundproof = $8, fsmonkingarea = $9, ftelevison = $10, fac = $11, fbathroom = $12, fcoffemaker = $13, updatedat = $14 WHERE id = $15 RETURNING id, roomdescription, fwifi, fsoundsystem, fprojector, fscreenprojector, fchairs, ftables, fsoundproof, fsmonkingarea, ftelevison, fac, fbathroom, fcoffemaker,createdat, updatedat`, facilitie.RoomDescription, facilitie.Fwifi, facilitie.FsoundSystem, facilitie.Fprojector, facilitie.FscreenProjector, facilitie.Fchairs, facilitie.Ftables, facilitie.FsoundProof, facilitie.FsmonkingArea, facilitie.Ftelevison, facilitie.FAc, facilitie.Fbathroom, facilitie.FcoffeMaker, time.Now(), facilitie.Id).Scan(
		&facilitie.Id,
		&facilitie.RoomDescription,
		&facilitie.Fwifi,
		&facilitie.FsoundSystem,
		&facilitie.Fprojector,
		&facilitie.FscreenProjector,
		&facilitie.Fchairs,
		&facilitie.Ftables,
		&facilitie.FsoundProof,
		&facilitie.FsmonkingArea,
		&facilitie.Ftelevison,
		&facilitie.FAc,
		&facilitie.Fbathroom,
		&facilitie.FcoffeMaker,
		&facilitie.CreatedAt,
		&facilitie.UpdatedAt,
	)

	room.Facility = facilitie

	return room, err
}

// Delete implements RoomRepository.
func (r *roomRepository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM rooms WHERE id = $1`, id)
	if err != nil {
		panic(err)
	}
	return err
}

// Get implements RoomRepository.
func (r *roomRepository) Get(id string) (model.Room, error) {
	var room model.Room
	err := r.db.QueryRow(`SELECT r.id, r.roomtype, r.capacity, f.id, f.roomdescription, f.fwifi, f.fsoundsystem, f.fprojector, f.fscreenprojector, f.fchairs, f.ftables, f.fsoundproof, f.fsmonkingarea, f.ftelevison, f.fac, f.fbathroom, f.fcoffemaker, f.createdat, f.updatedat, r.status, r.createdat, r.updatedat FROM rooms AS r JOIN facilities AS f ON f.id = r.facilities WHERE r.id = $1;`, id).Scan(
		&room.Id,
		&room.RoomType,
		&room.MaxCapacity,
		&room.Facility.Id,
		&room.Facility.RoomDescription,
		&room.Facility.Fwifi,
		&room.Facility.FsoundSystem,
		&room.Facility.Fprojector,
		&room.Facility.FscreenProjector,
		&room.Facility.Fchairs,
		&room.Facility.Ftables,
		&room.Facility.FsoundProof,
		&room.Facility.FsmonkingArea,
		&room.Facility.Ftelevison,
		&room.Facility.FAc,
		&room.Facility.Fbathroom,
		&room.Facility.FcoffeMaker,
		&room.Facility.CreatedAt,
		&room.Facility.UpdatedAt,
		&room.Status,
		&room.CreatedAt,
		&room.UpdatedAt,
	)
	if err != nil {
		return model.Room{}, err
	}

	return room, err
}

func (r *roomRepository) Create(payload model.Room) (model.Room, error) {

	var room model.Room

	// var facility []model.RoomFacility
	var roomFacility model.RoomFacility
	err := r.db.QueryRow(`INSERT INTO facilities (roomdescription, fwifi, fsoundsystem, fprojector, fscreenprojector, fchairs, ftables, fsoundproof, fsmonkingarea, ftelevison, fac, fbathroom, fcoffemaker, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id, roomdescription, fwifi, fsoundsystem, fprojector, fscreenprojector, fchairs, ftables, fsoundproof, fsmonkingarea, ftelevison, fac, fbathroom, fcoffemaker,createdat, updatedat`, payload.Facility.RoomDescription, payload.Facility.Fwifi, payload.Facility.FsoundSystem, payload.Facility.Fprojector, payload.Facility.FscreenProjector, payload.Facility.Fchairs, payload.Facility.Ftables, payload.Facility.FsoundProof, payload.Facility.FsmonkingArea, payload.Facility.Ftelevison, payload.Facility.FAc, payload.Facility.Fbathroom, payload.Facility.FcoffeMaker, time.Now()).Scan(
		&roomFacility.Id,
		&roomFacility.RoomDescription,
		&roomFacility.Fwifi,
		&roomFacility.FsoundSystem,
		&roomFacility.Fprojector,
		&roomFacility.FscreenProjector,
		&roomFacility.Fchairs,
		&roomFacility.Ftables,
		&roomFacility.FsoundProof,
		&roomFacility.FsmonkingArea,
		&roomFacility.Ftelevison,
		&roomFacility.FAc,
		&roomFacility.Fbathroom,
		&roomFacility.FcoffeMaker,
		&roomFacility.CreatedAt,
		&roomFacility.UpdatedAt,
	)
	if err != nil {
		return model.Room{}, err
	}

	room.Facility.Id = roomFacility.Id
	room.Facility = roomFacility

	err = r.db.QueryRow(`INSERT INTO rooms (roomtype, capacity, facilities ,status, updatedat) VALUES ($1, $2, $3, $4, $5) RETURNING id, roomtype, capacity, status, createdat, updatedat`, payload.RoomType, payload.MaxCapacity, roomFacility.Id, payload.Status, time.Now()).Scan(
		&room.Id, &room.RoomType, &room.MaxCapacity, &room.Status, &room.CreatedAt, &room.UpdatedAt)
	if err != nil {
		return model.Room{}, err
	}

	return room, err
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}
