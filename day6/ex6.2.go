package main

import "fmt"

func main() {
	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)
	fmt.Println(a == 9.0)
}
