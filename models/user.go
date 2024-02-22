package models

import "database/sql"

type User struct {
	Id       int64          `db:"ID" json:"id"`
	Username string         `db:"Username" json:"username"`
	Password string         `db:"Password" json:"password"`
	Level    sql.NullString `db:"Level" json:"level"`
	// Token    string `db:"Token" json:"token"`
}
