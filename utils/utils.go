package utils

import (
	"math"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(bytes), err
}

func CheckPasswordHash(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}

func GetCurrentTime() time.Time {
	return time.Now()
}

func ArredondarParaDuasCasas(valor float64) float64 {
	return math.Round(valor*100) / 100
}
