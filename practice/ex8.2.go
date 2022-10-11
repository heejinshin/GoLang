package main

import "fmt"

func main() {
	const pi1 float64 = 3.141592653589793228 // 상수
	var pi2 float64 = 3.141592653589793228 // 변수

	// pi1 = 3
	pi2 = 4

	fmt.Printf("원주율: %f\n", pi1)
	fmt.Printf("원주율: %f\n", pi2)
}