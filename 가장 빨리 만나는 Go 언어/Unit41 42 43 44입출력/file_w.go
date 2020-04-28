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

	// Fprint
	file1, _ := os.Create("hello1.txt")
	defer file1.Close()
	fmt.Fprint(file1, 1, 1.1, "Hello, world")

	// Fprintln
	file2, _ := os.Create("hello2.txt")
	defer file2.Close()
	fmt.Fprintln(file2, 1, 1.1, "Hello, world")

	// Fprintf
	file3, _ := os.Create("hello3.txt")
	defer file3.Close()
	fmt.Fprintf(file3, "%d,%f,%s", 1, 1.1, "Hello, world")
}
