package hutils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func EncodePassword(password, salt string) string {
	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha256.New)
	return hex.EncodeToString(dk)
}
