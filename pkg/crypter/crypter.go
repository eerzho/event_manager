package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

type Crypter struct {
	key string
}

func New(key string) *Crypter {
	return &Crypter{key: key}
}

func (c *Crypter) Encrypt(decrypted string) string {
	const op = "crypter::Encrypt"

	decodedKey, err := base64.URLEncoding.DecodeString(c.key)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return ""
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return ""
	}

	plaintext := []byte(decrypted)
	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Printf("%s: %v\n", op, err)
		return ""
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(cipherText)
}

func (c *Crypter) Decrypt(encrypted string) string {
	const op = "crypter::Decrypt"

	decodedKey, err := base64.URLEncoding.DecodeString(c.key)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return ""
	}

	block, err := aes.NewCipher(decodedKey)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return ""
	}

	cipherText, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		log.Printf("%s: %v\n", op, err)
		return ""
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
