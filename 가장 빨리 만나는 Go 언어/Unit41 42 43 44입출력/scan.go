/*
func Scan(a ...interface{}) (n int, err error): 콘솔에서 공백, 새 줄로 구분하여 입력을 받음
func Scanln(a ...interface{}) (n int, err error): 콘솔에서 공백으로 구분하여 입력을 받음
func Scanf(format string, a ...interface{}) (n int, err error): 콘솔에서 형식을 지정하여 입력을 받음
*/

package main

import "fmt"

func main() {
	//var s1, s2 string
	////n, _ := fmt.Scan(&s1, &s2) // fmt.Scan 함수의 두 번째 리턴값은 생략
	//n, _ := fmt.Scanln(&s1, &s2)
	//fmt.Println("입력 개수:", n)
	//fmt.Println(s1, s2)

	var num1, num2 int
	n, _ := fmt.Scanf("%d,%d", &num1, &num2) // 정수형으로 형식을 지정하여 입력을 받음
	fmt.Println("입력 개수:", n)
	fmt.Println(num1, num2)

}
