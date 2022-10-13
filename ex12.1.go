// 배열의 개념과 사용법 
	// 최근 5일 온도 데이터를 저장하는 배열 변수 t를 선언한다. 
	// var t [5]float64
package main 

import "fmt"

func main() {   // 배열 t 선언, 키워드 var 사용, 변수 name 't'; 특징: 변수를 하나만 할당해도 된다 
	// var 변수명 [요소개수]타입 ; 타입앞에 요소의 개수가 온다.
	// 배열을 선언하고 초기화 한다. 중괄호 {} 를 이용하여 각 요소의 값을 초기화.  
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	//배열 길이에 따른 반복 횟수 지정. ; 5
	for i := 0; i < 5; i++ {
		fmt.Println(t[i])  // 인덱스 기법으로 출력 
	}
}