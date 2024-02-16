package models

type User struct {
	Id       int64  `db:"ID" json:"id"`
	Username string `db:"Username" json:"username"`
	Password string `db:"Password" json:"password"`
	Level    string `db:"Level" json:"level"`
	// Token    string `db:"Token" json:"token"`
}
