package main

import (
	"fmt"
	"reflect"
)

/*
	리플렉션 : 실행 시점(Runtime, 런타임)에 인터페이스나 구조체 등의 타입 정보를 얻어내거나 결정하는 기능
*/

type Data struct {
	a, b int
}

func main() {
	var num int = 1
	fmt.Println(reflect.TypeOf(num))

	var s string = "hello"
	fmt.Println(reflect.TypeOf(s))

	var f float64 = 1.2
	fmt.Println(reflect.TypeOf(f))

	var data Data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data))

	var numArr1 [3]int
	fmt.Println(reflect.TypeOf(numArr1))

	var numArr2 [5]int
	fmt.Println(reflect.TypeOf(numArr2))

	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)
	fmt.Println(t.Name())                    // float64: 자료형 이름 출력
	fmt.Println(t.Size())                    // 8: 자료형 크기 출력
	fmt.Println(t.Kind() == reflect.Float64) // true: 자료형 종류를 알아내어 reflect.Float64와 비교
	fmt.Println(t.Kind() == reflect.Int64)   // false: 자료형 종류를 알아내어 reflect.Int64와 비교
	fmt.Println(v.Type())                    // float64: 값이 담긴 변수의 자료형 이름 출력
	fmt.Println(v.Kind() == reflect.Float64) // true: 값이 담긴 변수의 자료형 종류를 알아내어 reflect.Float64와 비교
	fmt.Println(v.Kind() == reflect.Int64)   // false: 값이 담긴 변수의 자료형 종류를 알아내어 reflect.Int64와 비교
	fmt.Println(v.Float())                   // 1.3: 값을 실수형으로 출력
}
