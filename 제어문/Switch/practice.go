package main 

import "fmt"

func getAge() int{   // getAge의 결괏값은 22 
	return 22
}

func main() {
	switch age := getAge(); age{  // 결괏값이 age, int로 비교
	case 20:
		fmt.Println("스물")
	case 21:
		fmt.Println("스물하나")
	case 22:
		fmt.Println("스물 둘")
	}	
}