package malfun

import (
    "fmt"
    "errors"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
)

func DECPT(key []byte, secure string) (decoded string, err error) {
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)
	if err != nil {
		fmt.Printf("FUCK1")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("FUCK2")
	}
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}