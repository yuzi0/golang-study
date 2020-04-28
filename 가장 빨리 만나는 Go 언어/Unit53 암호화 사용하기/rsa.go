package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

/*
func GenerateKey(random io.Reader, bits int) (priv *PrivateKey, err error): 개인 키와 공개키 생성
func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) (out []byte, err error): 평문을 공개 키로 암호화
func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) (out []byte, err error): 암호화된 데이터를 개인 키로 복호화
*/

func main() {
	// 개인키 생성
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 개인키에서 공개키 추출
	publicKey := &privateKey.PublicKey

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	ciphertext, err := rsa.EncryptPKCS1v15(
		rand.Reader,
		publicKey,
		[]byte(s),
	)

	fmt.Printf("%x\n", ciphertext)

	plaintext, err := rsa.DecryptPKCS1v15(
		rand.Reader,
		privateKey,
		ciphertext,
	)

	fmt.Println(string(plaintext))
}
