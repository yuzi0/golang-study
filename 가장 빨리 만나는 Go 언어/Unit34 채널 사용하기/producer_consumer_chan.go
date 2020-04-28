package main

import "fmt"

// 보내기 전용 채널 chan <- T
func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}
	c <- 100
	// 보내기 전용 채널에서 값을 받으려고 하면 컴파일 에러 발생
}

// 받기 전용 채널 <-chan T
func consumer(c <-chan int) {
	// range 키워도 또는 <- 채널 형식으로 값을 꺼낼 수 있다
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c)
	// 받기 전용 채널에서 값을 보내려고 하면 컴파일 에러 발생
}

func main() {
	c := make(chan int)

	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
