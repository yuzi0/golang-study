package main

import (
	"crypto/aes"
	"fmt"
)

/*
func NewCipher(key []byte) (cipher.Block, error): 대칭키 암호화 블록 생성
func (c *aesCipher) Encrypt(dst, src []byte): 평문을 AES 알고리즘으로 암호화
func (c *aesCipher) Decrypt(dst, src []byte): AES 알고리즘으로 암호화된 데이터를 평문으로 복호화
*/

func main() {
	key := "Hello, Key 12345"
	s := "Hello, world! 12"

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plaintext := make([]byte, len(s))
	block.Decrypt(plaintext, ciphertext)
	fmt.Println(string(plaintext))
}
