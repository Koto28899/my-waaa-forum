package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := GenerateRandomString(12)

	hashwdPwd, err := HashedPwd(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashwdPwd)

	err = CheckPwd(password, hashwdPwd)
	require.NoError(t, err)

	wrongPassword := GenerateRandomString(12)
	err = CheckPwd(wrongPassword, hashwdPwd)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashwdPwd2, err := HashedPwd(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashwdPwd2)
	require.NotEqual(t, hashwdPwd, hashwdPwd2)
}
