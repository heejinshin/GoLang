package main 
import "fmt"

func Add(a int, b int) int { // 입력 값, 반환 값 모두 정수
	return a+b
}

func main() {
	c:= Add(3, 6)
	fmt.Println(c)
}
