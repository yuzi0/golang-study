package main

import (
	"fmt"
	"time"
)

// 여러 채널을 손쉽게 사용할 수 있도록 select 분기문을 사용한다
// switch 분기문과 비슷하지만 select 키워드 뒤에 검사할 변수를
// 따로 지정하지 않으면 각 채널에 값이 들어오면 해당 case만 싱행된다

func main() {
	c1 := make(chan int)
	c2 := make(chan string)

	go func() {
		for {
			i := <-c1
			fmt.Println("c1 :", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for {
			c2 <- "Hello world"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			select {
			case c1 <- 10:
			case s := <-c2:
				fmt.Println("c2 :", s)
			case <-time.After(50 * time.Millisecond):
				fmt.Println("timeout")
			}
		}
	}()

	time.Sleep(10 * time.Second)

}
