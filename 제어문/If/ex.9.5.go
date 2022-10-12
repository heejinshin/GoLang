package main

import "fmt"

func IsThereRoze() bool {
	return true // 일단 있는 걸로 설정
}

func GetFriendsCount() int { // 같이간 친구 수, 이걸 뭐하러 반환을 하지
	return 3
}

func main() {
	price := 9000   // roze 가격 
    myprice := 9000 // 내가 가진돈 
	friends := 3

	if price >= 30000 {
		if friends >= 3 {
			fmt.Println("더치ㄱㄱ")
		} else {
			fmt.Println("음 고민좀")
		}
	} else if price >= 10000 && price < 30000 {
        if myprice >= 10000 {
            fmt.Println("혼자 먹음")
        } else {
            fmt.Println("더치해")
        }

	}else {
        fmt.Println("내가 쏨")
    }
}