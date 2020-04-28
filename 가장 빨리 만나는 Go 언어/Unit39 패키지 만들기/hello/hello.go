package main

import (
	"calc"
	"fmt"
)

func main() {
	// calc 패키지의 Sum 함수 사용
	fmt.Println(calc.Sum(1, 2))
}

// calc 패키지를 컴파일하여 라이브러리로 만들려면 cmd에서 src/calc 경로에서 go install 명령을 입력한다.
// 또는 go install calc 명령을 입력한다.

// 그럼 pkg/운영체제_아키텍쳐에 calc.a 라이브러리 파일이 생성된다.
