// 중첩 

package main 

import "fmt"

func main() {


	for dan := 2; dan <=9; dan++{
		for a := 1; a <=9 ; a++{
			fmt.Printf("%d * %d = %d\n", dan, a, dan * a)
		}
	}
	fmt.Print("for문이 종료됐습니다.")
}