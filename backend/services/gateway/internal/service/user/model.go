package user

import "time"

type User struct {
	Id           int
	Email        string
	Password     string
	RegisterDate time.Time
}
