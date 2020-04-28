package main

import "fmt"

func main() {
	// int형 채널 생성
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			// 채널에 값 보냄
			c <- i
		}
		// 채널 닫음
		close(c)
	}()

	// range를 사용해서 채널에서 값을 꺼냄
	for i := range c {
		// 꺼낸 값을 출력
		fmt.Println(i)
	}
}

/*
	이미 닫힌 채널에 값을 보내면 패닉이 발생한다
	채널을 닫으면 range 루프가 종료된다
	채널이 열려있고, 값이 들어오지 않는다면 range는 실행되지 않고 계속 대기한다.
	만약 다른 곳에서 채널에 값을 보냈다면(채널에 값이 들어오면) 그때부터 range가 계속 반복된다
	==> range는 채널에 값이 몇 개가 들어올지 모르기 때문에 값이 들어올 때마다 계속 꺼내기 위해 사용한다
*/
