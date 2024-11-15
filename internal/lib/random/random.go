package random

import (
	random "math/rand"
	"time"
)

func NewRandomString() string {
	r := random.New(random.NewSource(time.Now().UnixNano()))

	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 8)
	for i := range b {
		b[i] = chars[r.Intn(len(chars))]
	}

	return string(b)
}
