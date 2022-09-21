package main

import (
    "encoding/base64"
    "crypto/cipher"
    "crypto/aes"
	"crypto/rand"
	"flag"
    "fmt"
    "io"
    "log"
)

func encrypt(key []byte, message string) (encoded string, err error) {
	plainText := []byte(message)
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText), err
}

func main() {
	token := flag.String("t", "", "Dicord Bot Token")
	flag.Parse()
	key := []byte("AB1g4ssBuNnyJumPingUpTheHillBill")
    ciphertext, err := encrypt(key, *token)
    if err != nil {
        log.Fatal(err)
    }
	fmt.Printf("%s\n", ciphertext)
}