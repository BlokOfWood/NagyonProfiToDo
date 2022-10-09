package main

import (
	"crypto/sha256"
	"math/rand"
	"time"

	Server "github.com/CodeFoxHu/go-serverlib"
)

const hextable = "0123456789abcdef"
const SaltLength uint8 = 8

var saltRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func EncodedLen(n int) int { return n * 2 }

func Encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}

func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

func EncodePassword(password string, salt string) string {
	password = salt + password
	hash := sha256.Sum256([]byte(password))

	return EncodeToString(hash[:])
}

func GenerateSalt() string {
	b := make([]rune, SaltLength)
	for i := range b {
		b[i] = saltRunes[rand.Intn(len(saltRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetUserFromDB(ctx *Server.FoxContext) error {

	return nil
}
