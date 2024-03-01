package utils

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"
)

func Init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateRandomBool() bool {
	return rand.Intn(2) == 1
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func GenerateRandomInt63(lo int64, hi int64) int64 {
	return lo + rand.Int63n(hi-lo-1)
}

func GenerateRandomString(n int) string {
	var s strings.Builder
	l := len(alphabet)
	for i := 0; i < n; i++ {
		s.WriteByte(alphabet[rand.Intn(l)])
	}
	return s.String()
}

func GenerateRandomEmail() string {
	return (GenerateRandomString(20) + "@" + GenerateRandomString(10) + ".com")
}

func GenerateRandomStatement() sql.NullString {
	if GenerateRandomBool() {
		return sql.NullString{
			Valid:  true,
			String: GenerateRandomString(30),
		}
	} else {
		return sql.NullString{
			Valid: false,
		}
	}
}
