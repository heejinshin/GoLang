// 다음 예제의 결과를 쓰세요

package main 
import "fmt"

func ChangeArray(arr [5]int) {  // 입력값으로 int 5개짜리 arr를 받음 
	arr[3] = 3000    // 배열 arr의 4번째 값을 3000으로 바꿈 
}

func main() {  
	a := [5]int{1, 2, 3, 4, 5}

	ChangeArray(a)  // change~ 함수에 a를 넣으면 배열 a[3]은 3000으로 바뀔 것임. 
	fmt.Print(a[3])   // a는 현재 어떤 상태냐, {1, 2, 3, 3000, 5}
}

// 