// 곱해서 45가 나오는 수 찾는 프로그램 _ 플래그 변수 사용법 

package main 

import "fmt"


func main() {
	a := 1 
	b := 1
	found := false 
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a * b == 45 {
				found = true
				break
			}
		}
		if found {
			break 
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}