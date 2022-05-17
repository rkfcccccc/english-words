package verification

import "time"

type Entry struct {
	Id         string    `db:"id"`
	Email      string    `db:"email"`
	TypeId     int       `db:"type_id"`
	Code       int       `db:"code"`
	Attempts   int       `db:"attempts"`
	ExpireDate time.Time `db:"expire_time"`
}
