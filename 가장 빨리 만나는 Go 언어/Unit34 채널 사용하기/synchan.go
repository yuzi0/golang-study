package main

import (
	"fmt"
	"time"
)

func main() {
	// 동기채널 생성
	done := make(chan bool)
	count := 3

	go func() {
		for i := 0; i < count; i++ {
			done <- true
			fmt.Println("goroutine :", i)
			// 1초 대기
			time.Sleep(1 * time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		// 채널에 값이 들어올 때까지 대기, 값을 꺼냄
		<-done
		fmt.Println("main func :", i)
	}
}

// 고루틴 -> 메인 함수 -> 고루틴 -> 메인 함수
// 동기 채널은 보내는 쪽에서는 값을 받을 때까지 대기하고, 받는 쪽에서는 채널에 값이 들어올 때까지 대기
// ==> 동기 채널을 활용하면 고루틴의 코드 실행 순서를 제어할 수 있음
