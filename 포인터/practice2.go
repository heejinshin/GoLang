//Actor 구조체를 생성하고 주어진 인수를 이용해서 값을 초기화해서 변환하는 MakeActor() 함수를 완성하세요 . 

package main 

import "fmt"

type Actor struct {
	Name string
	HP int 
	Speed float64
}

func NewActor(name string, hp int, speed float64) *Actor {
	
}