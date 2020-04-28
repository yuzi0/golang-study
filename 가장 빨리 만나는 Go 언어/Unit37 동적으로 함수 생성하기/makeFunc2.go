package main

import (
	"fmt"
	"reflect"
)

// 매개 변수와 리번은 반드시 []reflect.Value 사용
func sum(args []reflect.Value) []reflect.Value {
	a, b := args[0], args[1]
	if a.Kind() != b.Kind() {
		fmt.Println("타입이 다릅니다.")
		return nil
	}

	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []reflect.Value{reflect.ValueOf(a.Int() + b.Int())}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []reflect.Value{reflect.ValueOf(a.Uint() + b.Uint())}
	case reflect.Float32, reflect.Float64:
		return []reflect.Value{reflect.ValueOf(a.Float() + b.Float())}
	case reflect.String:
		return []reflect.Value{reflect.ValueOf(a.String() + b.String())}
	default:
		return []reflect.Value{}
	}
}

func makeSum(fptr interface{}) {
	// 1. 인터페이스 주소를 넘긴 뒤 Elem으로 값 정보를 가져옴
	// 2. sum의 함수 정보를 생성
	// 3. 인터페이스의 값 정보인 fn에 sum의 함수 정보를 v를 설정하여 함수를 연결
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), sum)
	fn.Set(v)
}

func main() {
	var intSum func(int, int) int64
	var floatSum func(float32, float32) float64
	var stringSum func(string, string) string

	makeSum(&intSum)
	makeSum(&floatSum)
	makeSum(&stringSum)

	fmt.Println(intSum(1, 2))                   // 3
	fmt.Println(floatSum(2.1, 3.5))             // 5.599999904632568
	fmt.Println(stringSum("Hello, ", "world!")) // Hello, world!
}
