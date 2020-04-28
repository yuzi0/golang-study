package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
	Once를 사용하면 함수를 한 번만 실행할 수 있습니다.
	func (*Once) Do(f func()): 함수를 한 번만 실행
*/
func hello() {
	fmt.Println("Hello world")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	once := new(sync.Once) // Once 생성

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine : ", n)
			once.Do(hello)
		}(i)
	}
	fmt.Scanln()
}
