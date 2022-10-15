package main 

import "fmt"

fun main() {
	var a int = 500
	var p *int

	p = &a // a 의 메모리 주소를 변수 p의 값으로 대입(복사)

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100
	fmt.Printf("a의 값: %d\n", a)
}