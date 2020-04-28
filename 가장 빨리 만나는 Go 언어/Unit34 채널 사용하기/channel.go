package main

import "fmt"

func sum(a int, b int, c chan int) {
	// 채널에 a와 b의 합을 보냄
	c <- a + b
}

func main() {
	// 채널을 사용하기 전에는 반드시 make 함수로 공간을 할당해야 한다
	// 이렇게 생성하면 동기 채널(synchronous channel)이 생성된다

	// int형 채널 생성
	c := make(chan int)
	// sum을 고루틴으로 실행한 뒤, 채널을 매개변수로 넘겨줌
	go sum(1, 2, c)
	// 채널에서 값을 꺼낸 뒤, n에 대입
	n := <-c

	fmt.Println(n)
}

/*
채널(channel) : 고루틴끼리 데이터를 주고 받고, 실행 흐름을 제어하는 기능
- 모든 타입을 채널로 사용할 수 있음
- 채널 자체는 값이 아닌 레퍼런스 타입임

:=를 사용하지 않고, var 키워드로 채널을 선언하고 할당할 수 있다
chan int형 변수 선언
	var c chan int
	c = make(chan int)

채널 <- 값 : 채널로 값 보내기
변수 <- 채널 : 채널에서 값 가져오기
*/
