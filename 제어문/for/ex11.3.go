// 숫자를 입력받다가 짝수이면 프로그램을 종료한다 
// 반복문을 사용한다 

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("입력하세요.?뭘?")
		var number int
		_, err:= fmt.Scanln(&number)  // 한 줄 입력 받기 
		if err != nil { // 에러가 있으면
			fmt.Println(err, "숫자를 입력하세요.")
			//키보드 버퍼를 비웁니다. 
			stdin.ReadString('\n')
			continue
		}
		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)
		if number % 2 == 0 {
			break
		}
	}
	fmt.Println("for문이 종료됐습니다.")
}