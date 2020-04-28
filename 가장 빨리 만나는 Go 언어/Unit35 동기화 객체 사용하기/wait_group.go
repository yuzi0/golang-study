package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
대기 그룹 : 고루틴이 모두 끝날 때까지 기다릴 때 사용

func (wg *WaitGroup) Add(delta int): 대기 그룹에 고루틴 개수 추가
func (wg *WaitGroup) Done(): 고루틴이 끝났다는 것을 알려줄 때 사용
func (wg *WaitGroup) Wait(): 모든 고루틴이 끝날 때까지 기다림

- Add와 Done의 호출 횟수는 같아야 한다 ==> 그렇지 못할 경우 패닉 발생
- Done 함수는 defer와 함께 사용해서 지연 호출로도 사용할 수 있다
- defer : 특정 문자 혹은 함수를 나중에 (defer를 호출하는 함수가 리턴하기 직전에) 실행하는 함.
*/

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 대기 그룹 생성
	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		// 대기 그룹 수 추가
		wg.Add(1)
		go func(n int) {
			// defer wg.Done()
			fmt.Println(n)
			// 고루틴이 끝난다는 것을 알려줌
			wg.Done()
		}(i)
	}

	// 모든 고루틴이 끝날 때까지 기다림
	wg.Wait()
	fmt.Println("the end")
}
