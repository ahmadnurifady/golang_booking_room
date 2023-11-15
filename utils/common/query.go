package common

const (
	//!COMMON FOR USER
	CreateUser = `INSERT INTO users (name,divisi,jabatan,email,password,role, updatedat) VALUES ($1,$2,$3,$4,$5,$6,$7)
                  RETURNING id,name,divisi,jabatan,email,role,createdat, updatedat`
	GetUserById = `SELECT id,name,divisi,jabatan,email,role,createdat,updatedat FROM users WHERE id = $1`
)
