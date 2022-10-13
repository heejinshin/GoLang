// 배열은 그 길이가 정해져있기 때문에, 변숫값을 배열의 길이로 지정할 수 없다. 

package main 
import "fmt"

const Y int = 3

func main() {
	x := 5
	a := [x]int{1, 2, 3, 4, 5} // invalid array length x

	b := [Y]int{1, 2, 3}

	var c [10]int

	fmt.Print(x, a, b, c)
}