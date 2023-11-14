package repository

import (
	"database/sql"
	"final-project-booking-room/model"
	"final-project-booking-room/utils/common"
	"time"
)

type UserRepository interface {
	GetById(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
}

type userRepository struct {
	db *sql.DB
}

// !MENCARI USER BERDASARKAN ID
func (u *userRepository) GetById(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.GetUserById, id).
		Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// !MEMBUAT USER BARU
func (u *userRepository) Create(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.CreateUser, payload.Name, payload.Divisi, payload.Jabatan,
		payload.Email, payload.Password, payload.Role, time.Now()).
		Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
