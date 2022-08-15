package user

import (
	"time"

	pb "github.com/rkfcccccc/english_words/proto/user"
)

func transformFromGRPC(user *pb.User) *User {
	if user == nil {
		return nil
	}

	return &User{
		Id:           int(user.Id),
		Email:        user.Email,
		Password:     user.Password,
		RegisterDate: time.Unix(user.RegisterDate, 0),
	}
}
