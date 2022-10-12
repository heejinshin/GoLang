//fallthrough 사용해보기 // 사용 지양함 

package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough         // 여기서 case가 맞든 아니든 다음꺼 이어서 수행하게끔
	case 4:
		fmt.Println("a == 4")
	case 5:
		fmt.Println("a == 5")
	case 6:
		fmt.Println("a == 6")
	}
}