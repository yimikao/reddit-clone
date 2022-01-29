package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	l := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()

}
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomOwner() string {
	return RandomString(6)
}
