package main 
import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	} // b가 0 이 아니면 
	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}

// Divide() 함수 정의 , int 타입 a, b 를 매개변수로 받고 int 타입과 bool 타입 반환 
// (a int, b int) 같이 매개변수 타입이 같으면 (a. b int) 로 간단히 표현 

// 나눗셈 제수가 0이면 0과 false 를 반환 
// 제수가 0이 아닐 때는 나눗셈 결과와 true 반환 

// Divide 함수 호출, 그 결과를 c 와 success 로 받음 