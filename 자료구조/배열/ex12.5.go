// 배열의 복사 

package main 
import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {   // 배열 a의 요소 출력 (배열 a 순회)
		fmt.Printf("a[%d] = %d\n,", i, v)
	}

	fmt.Println()
	for i, v := range b {    // 배열 b 순회
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a   // 배열 a를 변수 b에 복사 

	fmt.Println() 
	for i, v := range b {   // range 가 배열 b에 대해서 인덱스 i와 값 v를 반환함 
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}