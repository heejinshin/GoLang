package main

import "fmt"

func getPrice() int{
	return 31
}

func main() {
	switch price := getPrice(); true{  // 조건 age랑 비교할 것임, 아마 default로 true일듯 
	case price >= 30:
		fmt.Println(" So expensive")
	case price > 20 :  // 조건 1에서 아닌거니까 30보단 아래 
		fmt.Println("soso")
	case price > 10:
		fmt.Println("not bad")
	}
}