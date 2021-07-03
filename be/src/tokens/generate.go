package tokens

import (
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"strings"
)

const (
	randchar = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJklmnopqrstuvwxyz1234567890"
)

func GenerateToken(n uint) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func Hash(v []byte) []byte {
	sha := sha1.New()
	sha.Write(v)
	return sha.Sum(nil)
}

func RandomID(n int) string {
	sb := &strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(randchar[rand.Intn(len(randchar)-1)])
	}
	return sb.String()
}
