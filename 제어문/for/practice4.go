// 다음 출력 결과가 나오도록 9까지 홀수의 제곱값을 출력하는 프로그램을 작성하세요 
	// 후처리는 꼭 ++나 --를 쓰지 않고 다른 연산자를 써도 됩니다. 
	// 1 * 1 = 1
	// 3 * 3 = 9 
	// 5 * 5 = 25  ......

package main 

import "fmt"

func main() {
	for a := 1; a <=9 ; a += 2{
		b := a * a
		fmt.Printf("%d * %d = %d\n", a, a, b)
	}
}
