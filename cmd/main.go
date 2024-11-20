package main

import (
	"fmt"
	"github.com/tilthyx/cryptography/cmd/decrypt"
	"github.com/tilthyx/cryptography/cmd/encrypt"
)

func main() {
	enc, err := encrypt.Encrypt("LUCAS VIEIRA")
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
