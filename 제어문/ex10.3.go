package main

import "fmt"

var msg1 string = "정상출근주입니다"
var msg2 string = "리프레시주입니다"

func PrintMsg1() {
	fmt.Print(msg1)
}

func PrintMsg2() {
	fmt.Print(msg2)
}

func main() {
	week :=4 
	
	switch week {
	case 1:
	PrintMsg1()
	case 2: 
	PrintMsg1()
	case 3:
	PrintMsg1()
	case 4:
	PrintMsg2()
	}
}