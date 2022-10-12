// break를 사용하지 않아도 case 하나 실행 후 자동으로 빠져나감, 비교해보기 ;
// 대신에 fallthorough 사용함 

package main 
import "fmt"

func main() {
	a := 2

	switch a{
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("yes")
	case 3:
		fmt.Println("3")
		break     // 실행 하든말든 2에서 yes 하고 빠져나감 
	case 4: 
		fmt.Println("4")
	case 5: 
		fmt.Println("5")
		break
	}
}