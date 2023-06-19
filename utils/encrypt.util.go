package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

func generateRandomSalt(saltSize int) []byte {
	salt := make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}

	return salt
}

func Encrypt(password string) ([]byte, []byte) {
	salt := generateRandomSalt(384)
	hash := pbkdf2.Key([]byte(password), salt, 2048, 768, sha512.New)

	return salt, hash
}

func Verify(password string, hash []byte, salt []byte) bool {
	return string(hash) == string(pbkdf2.Key([]byte(password), salt, 2048, 768, sha512.New))
}

func EncodingBytes(str []byte) string {
	return hex.EncodeToString(str)
}

func DecodingString(str string) []byte {
	decode, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return decode
}
