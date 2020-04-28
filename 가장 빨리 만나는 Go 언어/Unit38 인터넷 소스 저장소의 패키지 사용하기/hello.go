package main

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("olleh"))
}

/*
	명령 프롬프트에서 명령어를 통해 import로 설정한 패키지 받아오기

	1) GitHub 주소에서 패키지를 자동으로 받아온다
	~$ cd $GOPATH/src/hello
	~/hello_project/src/hello$ go get

	2) 주소를 입력해서 수동적으로 패키지를 받아올수도 있다
	$ go get github.com/golang/example/stringutil
*/
