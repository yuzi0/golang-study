package main

import (
	"fmt"
	"os"
)

/*
- os 패키지
func Create(name string) (file *File, err error): 기존 파일을 열거나 새 파일을 생성
func Open(name string) (file *File, err error): 기존 파일을 열기
func (f *File) Close() error: 열린 파일을 닫음

- fmt 패키지
func Fprint(w io.Writer, a ...interface{}) (n int, err error): 값을 그대로 문자열로 만든 뒤 파일에 저장
func Fprintln(w io.Writer, a ...interface{}) (n int, err error): 값을 그대로 문자열로 만든 뒤 문자열 끝에 개행 문자(\n)를 붙이고 파일에 저장
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error): 형식을 지정하여 파일에 저장
func Fscan(r io.Reader, a ...interface{}) (n int, err error): 파일을 읽은 뒤 공백, 개행 문자로 구분된 문자열에서 입력을 받음
func Fscanln(r io.Reader, a ...interface{}) (n int, err error): 파일을 읽은 뒤 공백으로 구분된 문자열에서 입력을 받음
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error): 파일을 읽은 뒤 문자열에서 형식을 지정하여 입력을 받음
*/

func main() {
	var num1 int
	var num2 float32
	var s string

	// Fscan
	file1, _ := os.Open("hello1.txt")
	defer file1.Close()
	n, _ := fmt.Fscan(file1, &num1, &num2, &s)
	fmt.Println("입력 개수 :", n)
	fmt.Println(num1, num2, s)

	// Fscanln
	file2, _ := os.Open("hello2.txt")
	defer file2.Close()
	n, _ = fmt.Fscanln(file2, &num1, &num2, &s)
	fmt.Println("입력 개수 :", n)
	fmt.Println(num1, num2, s)

	// Fscanf
	file3, _ := os.Open("hello3.txt")
	defer file3.Close()
	n, _ = fmt.Fscanf(file3, "%d,%f,%s", &num1, &num2, &s)
	fmt.Println("입력 개수 :", n)
	fmt.Println(num1, num2, s)
}
