package user

import "time"

type User struct {
	Id           int       `db:"id"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	RegisterDate time.Time `db:"register_date"`
}
