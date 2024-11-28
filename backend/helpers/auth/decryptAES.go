package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"os"
)

func DecryptCredentials(filepath, key string)[]byte{
	ciphher, err := os.ReadFile(filepath)
	if err != nil {
        log.Fatal(err)
    }

	bytekey,err := hex.DecodeString(key)
	if err != nil {
        log.Fatal(err)
    }

	ciph, err := aes.NewCipher(bytekey)
	if err != nil {
        log.Fatal(err)
    }

	gcm, err := cipher.NewGCM(ciph)
	if err != nil {
        log.Fatal(err)
    }

	nonceSize := gcm.NonceSize()

	if len(ciphher) < nonceSize {
        log.Fatal(err)
    }

	nonce, ciphertext := ciphher[:nonceSize], ciphher[nonceSize:]
    credentials, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        log.Fatal(err)
    }
	return credentials
}