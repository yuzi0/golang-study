package main

import "fmt"

func hello() {
	fmt.Println("Hello world!")
}

func main() {
	// 함수를 고루틴으로 실행
	go hello()
	// main 함수가 종료되지 않도록 대기
	fmt.Scanln()
}

/*
고루틴 사용하기
- 고루틴 : 함수를 동시에 실행시키는 기능
- 스레드보다 운영체제의 리소스를 적게 사용하므로 많은 수의 고루틴을 쉽게 생성할 수 있음
- 호출할 함수 앞에 go 키워드를 붙이면 해당 함수는 고루틴으로 실행됨

*/
