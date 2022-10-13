// 함수를 만들어서 곱해서 45 되는 값 찾기 

package main 

import "fmt"

func find45(a int)(int bool) {
	for b :=1; b <=9; b++ {
		if a*b == 45 {
			return b, true
		}
	}
	return 0, false
}

func main() {
	a := 1
	b := 1

	for ; a <=9; a++ {
		var found bool 
		if b, found = find45(a); found [
			break 
		]
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
