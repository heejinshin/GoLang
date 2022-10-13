// for문이 복잡해질수록 플래그 변수 사용이 어려워진다. _ 레이블 사용법 
package main 

import "fmt"


func main() {
	a := 1
	b := 1


	OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b )
}