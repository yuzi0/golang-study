package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	s := "Hello World"

	// 코드를 수정하여 변수 i를 클로저의 매개변수로 넘기지 않고
	// fmt.Println으로 바로 출력하는 경우
	for i := 0; i < 100; i++ {
		// 클로저로 고루틴 실행하기
		go func() {
			fmt.Println(s, i)
		}()

		/*
			go func(n int) {
				fmt.Println(s, n)
			}(i)
		*/
	}
	fmt.Scanln()
}
