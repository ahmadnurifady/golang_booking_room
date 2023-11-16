package repository

import (
	"database/sql"
<<<<<<< HEAD
	"final-project-booking-room/model"
	"final-project-booking-room/utils/common"
=======
	"final-project/model"
	"final-project/utils/common"
>>>>>>> master
	"time"
)

type UserRepository interface {
	GetById(id string) (model.User, error)
	Create(payload model.User) (model.User, error)
<<<<<<< HEAD
=======
	UpdateUserById(id string, payload model.User) (model.User, error)
	DeleteUserById(id string) (model.User, error)
	GetAllUser() ([]model.User, error)
>>>>>>> master
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

<<<<<<< HEAD
=======
// !MENGUPDATE USER BERDASARKAN ID
func (u *userRepository) UpdateUserById(id string, payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(common.UpdateUser, payload.Name, payload.Divisi, payload.Jabatan,
		payload.Email, payload.Password, payload.Role, time.Now(), id).
		Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.UpdatedAt)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// !MENGHAPUS USER BERDASARKAN ID
func (u *userRepository) DeleteUserById(id string) (model.User, error) {
	_, err := u.db.Exec(common.DeleteUser, id)
	if err != nil {
		return model.User{}, err
	}

	return model.User{}, nil
}

// !MENAMPILKAN SEMUA USER
func (u *userRepository) GetAllUser() ([]model.User, error) {
	var users []model.User

	rows, err := u.db.Query(common.GetAllUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.Id, &user.Name, &user.Divisi, &user.Jabatan, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

>>>>>>> master
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}
