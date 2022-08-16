package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rkfcccccc/english_words/shared_pkg/cache"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
	ErrRefreshMiss  = errors.New("no user with that refresh token")
)

const refreshTTL = time.Hour * 24 * 30 * 6
const tokenTTL = time.Minute * 5

type UserClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type Helper struct {
	signingKey []byte
	cache      cache.Repository
}

func NewHelper(signingKey string, repo cache.Repository) *Helper {
	return &Helper{[]byte(signingKey), repo}
}
