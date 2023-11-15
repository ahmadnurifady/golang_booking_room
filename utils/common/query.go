package common

const (
	//!COMMON FOR USER
	CreateUser = `INSERT INTO users (name,divisi,jabatan,email,password,role, updatedat) VALUES ($1,$2,$3,$4,$5,$6,$7)
                  RETURNING id,name,divisi,jabatan,email,role,createdat, updatedat`
	GetUserById = `SELECT id,name,divisi,jabatan,email,role,createdat,updatedat FROM users WHERE id = $1`
	UpdateUser  = `UPDATE users SET name = $1, divisi = $2, jabatan = $3,
				email = $4, password = $5, role = $6, updatedat = $7 WHERE id = $8
				RETURNING id,name,divisi,jabatan,email,role,updatedat`
	DeleteUser = `DELETE FROM users WHERE id = $1`
	GetAllUser = `SELECT id,name,divisi,jabatan,email,role,createdat,updatedat FROM users`
)
