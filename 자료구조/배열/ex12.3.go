package main 
import "fmt"

// 배열 순회 첫 번째 방법 _ for 문 사용 ; len(arr)
// 배열 요소 읽기, 쓰기 
// 배열 변수에 '[접근하고자 하는 요소의 인덱스]'를 적어 값을 읽고 쓴다. 

func main() {
	nums := [...]int{10, 20, 30, 40, 50}

	nums[2] = 300   // 인덱스 2에 해당하는 값을 새로대입하여 바꾼다. 

	for i := 0; i < len(nums); i++ {    // leng(배열) ; 배열의 길이 반환 
		fmt.Println(nums[i])     // 인덱스를 배열의 길이-1 만큼 설정한 후 대입
	}
}