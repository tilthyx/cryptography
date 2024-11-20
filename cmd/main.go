package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/tilthyx/cryptography/cmd/decrypt"
	"github.com/tilthyx/cryptography/cmd/encrypt"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	Secret := os.Getenv("PASSWORD")

	enc, err := encrypt.Encrypt(Secret)
	if err != nil {
		panic(err)
	}

	fmt.Printf("String encrypted: %s\n", enc)

	dec, decErr := decrypt.Decrypt(enc)
	if decErr != nil {
		panic(decErr)
	}

	fmt.Printf("String decrypted: %s\n", dec)

}
