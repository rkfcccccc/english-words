package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rkfcccccc/english_words/services/gateway/pkg/cache"
)

var (
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
)

type UserClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

type Helper struct {
	signingKey string
	tokenTTL   time.Duration
	cache      cache.Repository
}

func NewHelper(signingKey string, tokenTTL time.Duration, repo cache.Repository) *Helper {
	return &Helper{signingKey, tokenTTL, repo}
}
