package main

import (
	"fmt"
	"reflect"
)

// 매개 변수와 리번은 반드시 []reflect.Value 사용
func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello world!")
	return nil
}

func main() {
	// 함수를 담을 변수 선언
	var hello func()
	// hello의 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴
	fn := reflect.ValueOf(&hello).Elem()
	// h의 함수 정보를 생성
	v := reflect.MakeFunc(fn.Type(), h)
	// hello의 값 정보인 fn에 h의 함수 정보를 v를 설정하여 함수를 연결
	fn.Set(v)

	hello()
}
