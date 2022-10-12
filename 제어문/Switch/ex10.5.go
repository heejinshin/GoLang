// 비굣값을 true로 설정하여 결괏값이 true가 나오는 경우를 실행하도록 만듦. 
// Switch로 if문처럼 true 가 되는 조건문을 검사하기 

package main

import "fmt"
func main() {
	temp := 40

	switch true {  // 비굣값을 적지 않을 때 true 사용 가능, default 값
	case temp > 30 :
		fmt.Println("아직은 여름같아요")
	case temp < 20 && temp > 10 : 
		fmt.Println("겉옷 꼭 챙겨다니세요")
	case temp <= 10 : 
		fmt.Println("겨울이야")
	}
}
