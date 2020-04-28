package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*

func (rw *RWMutex) RLock(), func (rw *RWMutex) RUnlock(): 읽기 뮤텍스 잠금 및 잠금 해제
	읽기 락(Read Lock): 읽기 락끼리는 서로를 막지 않습니다.
	하지만 읽기 시도 중에 값이 바뀌면 안 되므로 쓰기 락은 막습니다.

func (rw *RWMutex) Lock(), func (rw *RWMutex) Unlock(): 쓰기 뮤텍스 잠금, 잠금 해제
	쓰기 락(Write Lock): 쓰기 시도 중에 다른 곳에서 이전 값을 읽으면 안 되고,
	다른 곳에서 값을 바꾸면 안 되므로 읽기, 쓰기 락 모두 막습니다.

*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data int = 0
	var rwMutex = new(sync.RWMutex)

	go func() { // 값을 쓰는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1                         // data에 값 쓰기
			fmt.Println("write   : ", data)   // data 값을 출력
			time.Sleep(10 * time.Millisecond) // 10 밀리초 대기
			rwMutex.Unlock()
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 1 : ", data) // data 값을 출력(읽기)
			time.Sleep(1 * time.Second)    // 1초 대기
			rwMutex.RUnlock()
		}
	}()

	go func() { // 값을 읽는 고루틴
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2 : ", data) // data 값을 출력(읽기)
			time.Sleep(2 * time.Second)    // 2초 대기
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(10 * time.Second) // 10초 동안 프로그램 실행
}
