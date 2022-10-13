package main 

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} // int형태 요소 5개 갖는 배열 a  

	for i, v := range a {   // 배열 a 를 순회해서 인덱스 i와 값 v 반환 
		a[i] = v * 2     
	}

	fmt.Println(a[2])   // a의 세번째 값에 접근하는 건데  
	
	// print문을 출력하기 전에 for 문이 배열을 순회하면서 i 번째 인덱스에다가 v의 2배 값을 넣어 변경시킴 
}

