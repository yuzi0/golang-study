package main

import "fmt"

func main() {
	c := make(chan int, 0)

	go func() {
		c <- 1
	}()

	// 채널이 닫혔는지 확인
	a, ok := <-c
	// 1 true : 채널은 열려 있고 1을 꺼냄
	fmt.Println(a, ok)

	close(c)

	// 채널이 닫혔는지 확인
	a, ok = <-c
	// 0 false : 채널은 닫혀 있음
	fmt.Println(a, ok)

}
