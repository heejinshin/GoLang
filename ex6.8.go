package main
import (
	"fmt"
	"math"
)

func equal (a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.3
	var c float64 = 0.2

	fmt.Printf("%0.18f == %0.18f : %v\n", a+b, equal(a+b, c))
}