// 계산 패키지
package calc

// 두 정수를 더함
func Sum(a, b int) int {
	return a + b
}

/*
	첫 글자를 영문 소문자로 지정하면 패키지 안에서만 사용할 수 있습니다. 즉 외부에서 사용할 수 없습니다.
	예) sum, max, hello

	첫 글자를 영문 대문자로 지정하면 외부에서 사용할 수 있습니다.
	예) Sum, Max, Hello
*/
