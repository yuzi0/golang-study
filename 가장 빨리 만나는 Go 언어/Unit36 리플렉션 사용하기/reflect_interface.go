package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))  // *int
	fmt.Println(reflect.ValueOf(a)) // <*int Value>
	//fmt.Println(reflect.ValueOf(a).Int())        // 런타임 에러
	fmt.Println(reflect.ValueOf(a).Elem())       // <int Value>
	fmt.Println(reflect.ValueOf(a).Elem().Int()) // 1

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))        // int
	fmt.Println(reflect.ValueOf(b))       // <int Value>
	fmt.Println(reflect.ValueOf(b).Int()) // 1
	//fmt.Println(reflect.ValueOf(b).Elem()) // 런타임 에러
}

/*
	포인터와 인터페이스의 값 가져오기

	포인터는 일반 변수와는 다르게 값을 가져오려면
	reflect.ValueOf 함수로 값 정보 reflect.Value를 얻어온 뒤 다시 Elem 함수로 값 정보를 가져와야 합니다.
	그리고 변수의 타입에 맞는 Int, Float, String 등의 함수로 값을 가져옵니다.

	여기서는 int 포인터 a의 값 정보에서 바로 Int 함수로 값을 가져오려면 런타임 에러가 발생합니다.
	따라서 Elem 함수로 포인터의 메모리에 저장된 실제 값 정보를 가져와야 합니다.

	빈 인터페이스 b에 1을 대입하면 타입 정보는 int이고 값 정보는 int입니다.
	따라서 인터페이스의 값을 가져오려면 변수 타입에 맞는 Int, Float, String 등의 함수를 사용하면 됩니다.
*/
