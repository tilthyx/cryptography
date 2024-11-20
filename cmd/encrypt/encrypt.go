package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(text string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	Secret := os.Getenv("SECRET")

	block, err := aes.NewCipher([]byte(Secret))
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}
