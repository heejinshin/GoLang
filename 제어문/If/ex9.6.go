// 초기문 ; 조건문 형태
// 어떤 함수를 실행하고 그 함수의 결과를 검사할 때 주로 사용
// 함수 성공 여부에 따라 다른 메시지를 출력
package main

import (
	"fmt"

	
)

func getMyAge() (int, bool) {
	return 40, false
}

func main() {
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println ("1) You are super young", age)
	}else if ok && age < 30 {
		fmt.Println("1) You are young", age)
	}else if ok { // age는 ? 
		fmt.Println("1) You are ??", age)  // 여기도 age는 나옴
	}else {   // 모든 조건이 false 일때 
		fmt.Println("1) It's error", age)
	}


	fmt.Println("2) The end")  // 초기문에 선언된 변수는 if 문을 벗어나면 끝난다. 
}