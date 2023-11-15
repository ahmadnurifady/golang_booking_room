package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Divisi    string    `json:"divisi"`
	Jabatan   string    `json:"jabatan"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u User) IsValidRole() bool {
	return u.Role == "admin" || u.Role == "employee" || u.Role == "GA"
}
