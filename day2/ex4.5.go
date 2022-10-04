package main

import "fmt"

func main() {
	var a int16 = 3456 // int16 은 2byte 정수
	var b int8 = int8(a)

	fmt.Println(a, b)
}
