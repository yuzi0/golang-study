package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// mutex : 여러 고루틴이 공유하는 데이터를 보호할 때 사용됨
//  func(m *Mutex) Lock() : 뮤텍스 잠금
//  func(m *Mutex) Inlock() : 뮤텍스 잠금 해제

/*
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{} // int형 슬라이스 생성

	go func() { // 고루틴에서
		for i := 0; i < 1000; i++ { // 1000번 반복하면서
			data = append(data, 1) // data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() { // 고루틴에서
		for i := 0; i < 1000; i++ { // 1000번 반복하면서
			data = append(data, 1) // data 슬라이스에 1을 추가

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기

	fmt.Println(len(data)) // data 슬라이스의 길이 출력
} // 2000이 출력 안 됨
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용

	var data = []int{}
	var mutex = new(sync.Mutex)

	go func() { // 고루틴에서
		for i := 0; i < 1000; i++ { // 1000번 반복하면서
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // data 슬라이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	go func() { // 고루틴에서
		for i := 0; i < 1000; i++ { // 1000번 반복하면서
			mutex.Lock()           // 뮤텍스 잠금, data 슬라이스 보호 시작
			data = append(data, 1) // data 슬라이스에 1을 추가
			mutex.Unlock()         // 뮤텍스 잠금 해제, data 슬라이스 보호 종료

			runtime.Gosched() // 다른 고루틴이 CPU를 사용할 수 있도록 양보
		}
	}()

	time.Sleep(2 * time.Second) // 2초 대기

	fmt.Println(len(data)) // data 슬라이스의 길이 출력
} // 2000 출력

// Lock, Unlock 함수는 반드시 짝을 맞추어주어야 하며
// 짝이 맞지 않으면 데드락(deadlock, 교착 상태)이 발생하므로 주의합니다.
