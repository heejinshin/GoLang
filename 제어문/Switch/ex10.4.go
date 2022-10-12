// 한번에 여러값을 비교
package main

import "fmt"

func main() {
	
	day := "fri"

	switch day {
	case "mon", "tue", "wed", "thu":	
		fmt.Println("일하세요")
	case "fri":
		fmt.Println("노세요")
	}
}