package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrTokenNotValidYet = errors.New("token not valid yet")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidToken     = errors.New("token is invalid")
)

type Payload struct {
	RoleId int64 `json:"roleId"`
	jwt.RegisteredClaims
}

func NewPayload(roleId int64, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		RoleId: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenId.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	now := time.Now()
	if payload.IssuedAt.After(now) {
		return ErrTokenNotValidYet
	}
	if payload.ExpiresAt.Before(now) {
		return ErrExpiredToken
	}
	return nil
}
