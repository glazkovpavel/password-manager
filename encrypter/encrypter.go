package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY environment variable not set")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGSM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, plainStr, nil) // Seal шифрует

}

func (enc *Encrypter) Decrypt(encrypterStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encrypterStr[:nonceSize], encrypterStr[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
