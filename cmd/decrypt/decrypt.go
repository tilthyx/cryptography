package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Decrypt(text string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	Secret := os.Getenv("SECRET")

	block, errCipher := aes.NewCipher([]byte(Secret))
	if errCipher != nil {
		return "", errCipher
	}

	ciphertext := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
