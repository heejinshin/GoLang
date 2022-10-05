package main

import(
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)   //선언대입문
	// newReader 출력값이 타입이 됨 

	var a int
	var b int 
	
	n, err := fmt.Scanln(&a, &b)
	if err != nil { //에러가 발생했다면 
		fmt.Println(err)
		stdin.ReadString('\n') //표준 입력값을 받아 읽어 오는데, 개행 문자가 나올때 까지 읽어 와라 
	} else { // a, b가 제대로 입력 됐다면 
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	}else {
		fmt.Println(a, b)
	}

	
}
