package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
	원자적 연산 : 더 이상 쪼갤 수 없는 연산
	따라서 여러 스레드(고루틴), CPU 코어에서 같은 변수(메모리)를 수정할 때 서로 영향을 받지 않고 안전하게 연산할 수 있습니다.
	보통 원자적 연산은 CPU의 명령어를 직접 사용하여 구현되어 있습니다.

	원자적 연산에는 메모리 주소와 수정할 값을 넣습니다.
	따라서 atomic.AddInt32(&data, 1)처럼 & (참조)를 사용하여 data 변수의 메모리 주소를 대입합니다.

Add 계열: 변수에 값을 더하고 결과를 리턴합니다.
CompareAndSwap 계열: 변수 A와 B를 비교하여 같으면 C를 대입합니다. 그리고 A와 B가 같으면 true, 다르면 false를 리턴합니다.
Load 계열: 변수에서 값을 가져옵니다.
Store 계열: 변수에 값을 저장합니다.
Swap 계열: 변수에 새 값을 대입하고, 이전 값을 리턴합니다.
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			//data++
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//data--
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(data)

}
