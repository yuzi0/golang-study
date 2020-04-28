package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

/*
	crypto/cipher 패키지에서 제공하는 암호화 운용 모드 함수
func NewCBCEncrypter(b Block, iv []byte) BlockMode: 암호화 블록과 초기화 벡터로 암호화 블록 모드 인스턴스 생성
func (x *cbcEncrypter) CryptBlocks(dst, src []byte): 암호화 블록 모드 인스턴스로 암호화
func NewCBCDecrypter(b Block, iv []byte) BlockMode: 암호화 블록과 초기화 벡터로 복호화 블록 모드 인스턴스 생성
func (x *cbcDecrypter) CryptBlocks(dst, src []byte): 복호화 블록 모드 인스턴스로 복호화

	io 패키지에서 제공하는 읽기 함수
func ReadFull(r Reader, buf []byte) (n int, err error): io.Reader에서 buf의 길이만큼 데이터를 읽음
*/

func encrypt(b cipher.Block, plaintext []byte) []byte {
	// 블록 크기의 배수가 되어야 하므로 블록 크기에서 모자란 부분을 채워준다
	if mod := len(plaintext) % aes.BlockSize; mod != 0 {
		padding := make([]byte, aes.BlockSize-mod)
		plaintext = append(plaintext, padding...)
	}

	// 초기화 벡터 공간(aes.BlockSize)만큼 더 생성한다
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// 부분 슬라이스로 초기화 벡터 공간을 가져온다
	iv := ciphertext[:aes.BlockSize]
	// 랜덤 값을 초기화 벡터에 넣어준다
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println(err)
		return nil
	}

	// 암호화 블록과 초기화 벡터를 넣어서 암호화 블록 모드 인스턴스 생성
	mode := cipher.NewCBCEncrypter(b, iv)
	// 암호화 블록 모드 인스턴스로 암호화
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext

}

func decrypt(b cipher.Block, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("암호화된 데이터의 길이는 블록 크리의 배스가 되어야 합니다.")
		return nil
	}

	// 부분 슬라이스로 초기화 벡터 공간 가져오기
	iv := ciphertext[:aes.BlockSize]
	// 부분 슬라이스로 암호화된 데이터 가져오기
	ciphertext = ciphertext[aes.BlockSize:]

	// 평문 데이터를 저장할 공간을 생성
	plaintext := make([]byte, len(ciphertext))
	// 암호화 블록과 초기화 벡터를 넣어서 복호화 블록 모드 인스턴스 생성
	mode := cipher.NewCBCDecrypter(b, iv)
	// 복호화 블록 모드 인스턴스로 복호화
	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext
}

func main() {
	key := "Hello Key 123456" // 16바이트

	s := `동해 물과 백두산이 마르고 닳도록
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려강산
대한 사람, 대한으로 길이 보전하세.`

	block, err := aes.NewCipher([]byte(key)) // AES 대칭키 암호화 블록 생성
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := encrypt(block, []byte(s)) // 평문을 AES 알고리즘으로 암호화
	fmt.Printf("%x\n", ciphertext)

	plaintext := decrypt(block, ciphertext) // AES 알고리즘 암호문을 평문으로 복호화
	fmt.Println(string(plaintext))
}
