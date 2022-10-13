package main 

import "fmt"

func main() {

}


// 이중 배열 선언 


// var a = [5]int{1, 2, 3, 4, 5}  // int 형태요소 5개를 가진 배열 변수 a 

var b = [2][5]int {   // 5개 짜리 int 묶음, "[]" 배열 
	{1, 2, 3, 4, 5},  // b[0] 초기화용          
	{6, 7, 8, 9, 10} // b[1] 초기화용          
}