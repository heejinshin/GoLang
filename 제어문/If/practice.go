package main

import "fmt"

func main() {
	score := 80   // 현재 스코어
	count := 31  // 경기수

	if count < 30 {
		fmt.Println(" 평가불가", score)
		if count < 30 && score >= 80 {  // 중첩조건문이 실행되려면 일단 첫번째 조건문이 true여야 
			fmt.Println("점수가 나와버렸네요. score is ", score)
		} 
		
	}else if count > 30 && score >= 80{  // 경기수가 적으면 
		count += 10  // 경기수를 늘리고
		fmt.Println(" 경기를 더 진행합니다. ")
		score += 20
		if score >= 100 {
			fmt.Println("We win!!!")
		} else {
			fmt.Print(" 경기를 종료합니다.")
		}

	} else {
		fmt.Println("점수를 확인하세요")
	}
}