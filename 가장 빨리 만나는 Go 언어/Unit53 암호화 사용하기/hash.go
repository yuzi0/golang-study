package main

import (
	"crypto/sha512"
	"fmt"
)

/*
func New() hash.Hash: SHA512 해시 인스턴스 생성
func Sum512(data []byte) [Size]byte: SHA512 해시를 계산하여 리턴
func (d *digest) Write(p []byte) (nn int, err error): 해시 인스턴스에 데이터 추가
func (d0 *digest) Sum(in []byte) []byte: 해시 인스턴스에 저장된 데이터의 SHA512 해시 값 추출
*/

func main() {
	s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1)

	sha := sha512.New()
	sha.Write([]byte("Hello, "))
	sha.Write([]byte("world!"))
	h2 := sha.Sum(nil)
	fmt.Printf("%x\n", h2)
}
