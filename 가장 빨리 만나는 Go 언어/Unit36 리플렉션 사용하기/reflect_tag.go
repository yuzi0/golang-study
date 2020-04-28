package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

func main() {
	p := Person{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

	test, ok := reflect.TypeOf(p).FieldByName("test")
	fmt.Println(ok, test.Tag.Get("tag1"), test.Tag.Get("tag2"))
}

/*
	리플렉션으로 구조체의 태그 가져오기
	- 구조체 필드의 태그는 `태그명:"내용"` 형식으로 지정한다
	- 태그는 문자열 형태이며 문자열 안에 " " (따옴표)가 포함되므로 ` ` (백쿼트)로 감싸준다
	- 여러 개를 지정할 때는 공백으로 구분해준다
*/
