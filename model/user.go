package model

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"user_name"`
	Password string `db:"password"`
}
