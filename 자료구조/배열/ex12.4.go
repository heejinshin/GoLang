// range 순회 

package main 

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 31.1, 32.0}

	for i, v := range t {
		fmt.Println(i, v)
	}
}