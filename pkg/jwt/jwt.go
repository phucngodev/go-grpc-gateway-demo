package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenInvalid = errors.New("jwt token is invalid")
	ErrParseToken   = errors.New("couldn't parse token")
)

type Jwt struct {
	key   []byte
	issue string
}

type Claims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

func New(key []byte, issue string) *Jwt {
	return &Jwt{
		key:   key,
		issue: issue,
	}
}

func (j *Jwt) CreateToken(id uint, expire time.Duration) (string, error) {
	claims := Claims{
		id, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			Issuer:    j.issue,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ts, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}

	return string(ts), nil
}

func (j *Jwt) ParseToken(ts string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(ts, &Claims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return j.key, nil
	})
	if err != nil {
		return nil, ErrParseToken
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrTokenInvalid
	}

	return claims, nil
}
