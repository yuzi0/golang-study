package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

/*
	풀은 객체(메모리)를 사용한 후 보관해두었다가 다시 사용하게 해주는 기능입니다.
	객체를 반복해서 할당하면 메모리 사용량도 늘어나고, 메모리를 해제해야 하는 가비지 컬렉터에게도 부담이 됩니다.
	즉, 풀은 일종의 캐시라고 할 수 있으며 메모리 할당과 해제 횟수를 줄여 성능을 높이고자 할 때 사용합니다.
	그리고 풀은 여러 고루틴에서 동시에 사용할 수 있습니다.

func (p *Pool) Get() interface{}: 풀에 보관된 객체를 가져옴
func (p *Pool) Put(x interface{}): 풀에 객체를 보관

	- 메모리 효율에 좋음
	- 수명 주기가 짧은 객체는 pool에 적합하지 않음
*/
type Data struct { // Data 구조체 정의
	tag    string // 풀 태그
	buffer []int  // 데이터 저장용 슬라이스
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 풀 할당
	pool := sync.Pool{
		// Get 함수를 사용했을 때 호출될 함수 정의
		New: func() interface{} {
			// 새 메모리 할당
			data := new(Data)
			// 태그 설정
			data.tag = 0
			// 슬라이스 공간 할당
			data.buffer = make([]int, 10)
			// 할당한 메모리(객체) 리턴
			return data
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			// pool에서 *Data 타입으로 데이터를 가져옴
			data := pool.Get().(*Data)
			for index := range data.buffer {
				data.buffer[index] = rand.Intn(100)
			}
			fmt.Println(data)
			data.tag = "used"
			// pool에 객체 보관
			pool.Put(data)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get().(*Data)
			n := 0
			for index := range data.buffer {
				data.buffer[index] = n
				n += 2
			}
			fmt.Println(data)
			data.tag = "used"
			pool.Put(data)
		}()
	}

	fmt.Scanln()
}
