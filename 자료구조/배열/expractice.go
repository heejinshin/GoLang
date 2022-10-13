package main 
import "fmt"

func main() {
	var a [5]int = [5]int {1, 2, 3, 4, 5}

	for _, v := range a {  // '_' : 인덱스 무효화 
		fmt.Print(v)    // 인덱스 출력을 생략할 수 있다. 
	}
}