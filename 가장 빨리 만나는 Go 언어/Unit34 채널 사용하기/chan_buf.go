package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	// 버퍼가 2개인 비동기 채널 생성
	done := make(chan bool, 0)
	count := 10

	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("고루틴 :", i)
			done <- true
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수 :", i)
	}
}

/*
	< 비동기 채널 >
	보내는 쪽에서 버퍼가 가득 차면, 실행을 멈추고 대기하며
	받는 쪽에서는 버퍼에 값이 없으면 대기한다
*/
