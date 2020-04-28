package main

import (
	"fmt"
	"runtime"
)

func main() {
	// CPU 개수를 구한 뒤 사용한 최대 CPU 개수 설정
	runtime.GOMAXPROCS(runtime.NumCPU())
	// 설정 값 출력
	fmt.Println(runtime.GOMAXPROCS(0))

	s := "Hello world"

	for i := 0; i < 100; i++ {
		// 익명 함수를 고루틴으로 실행(클로저)
		go func(n int) {
			// s와 매개변수로 받은 n 값 출력
			fmt.Println(s, n)
		}(i)
		// 반복문의 변수는 매개변수로 넘겨줌
	}
	fmt.Scanln()
}
