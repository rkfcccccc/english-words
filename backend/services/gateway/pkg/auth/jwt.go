package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (helper *Helper) ParseJWT(tokenString string) (*UserClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(helper.signingKey), nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, ErrTokenExpired
		}

		return nil, ErrInvalidToken
	} else if err != nil {
		return nil, fmt.Errorf("jwt.Parse: %v", err)
	}

	if token.Valid {
		claims, ok := token.Claims.(UserClaims)
		if !ok {
			return nil, ErrInvalidToken
		}

		return &claims, nil
	}

	return nil, fmt.Errorf("parse failed: %v", err)

}

func (helper *Helper) IssueJWT(user_id int) (string, error) {
	claims := UserClaims{
		user_id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(helper.signingKey)
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}

	return signed, nil
}
