package token

import (
	"backend/utils"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.GenerateRandomString(minSercetKeySize))
	require.NoError(t, err)

	roleId := utils.GenerateRandomInt63(0, 999999999)
	duration := time.Hour
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(roleId, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, roleId, payload.RoleId)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiresAt.Time, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(utils.GenerateRandomString(minSercetKeySize))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.GenerateRandomInt63(0, 9999999999), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token has invalid claims: token is expired")
	require.Nil(t, payload)
}

func TestInvalidJWTAlgNone(t *testing.T) {
	payload, err := NewPayload(utils.GenerateRandomInt63(0, 9999999999), time.Minute)
	require.NoError(t, err)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.GenerateRandomString(minSercetKeySize))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token is unverifiable: error while executing keyfunc: token is invalid")
	require.Nil(t, payload)
}
