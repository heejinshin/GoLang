package main 

import "fmt"

func main() {
	for i := 2; i < 10; i++ { // i가 2일 때
		if i == 3 || i == 6 {
			continue
		}
		for j := 1; j < 10; j++ {  // j가 9까지 돌아서 2단 출력
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
		}
		fmt.Println()
	}
		
}