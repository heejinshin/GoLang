// 각  방에 불을 켜는 프로그램 
// 비트플래그 연산 예시. set, reset, ison 이 주요 쓰이는 코드이므로 봐두기 

package main

import "fmt"

const (
	MasterRoom uint8 = 1 << iota // 0000 0001
	LivingRoom // 0000 0010
	BathRoom  // 0000 0100
	Garage  // 0000 1000
	SmallRoom // 0001 0000
)

// 불을 켠다 
func SetLight(rooms, room uint8) uint8 {  // 어떤 룸을 켤건지, uint8을 받아 결과로 uint8
	return rooms | room    // 1 있으면 1이되도록 
}

// 불을 끈다 
func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room   // 1인 비트만 끄도록 
}

// 불이 켜져있는지 
func IsLightOn(rooms, room uint8) bool {
	return rooms & room == room   // 해당하는 비트값이 1이다 
}

// Room 상태에 따라 메시지 출력 
func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 불을 켠다")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}
}


func main() {
	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom) // 각 룸에 불을 켜라 
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom)  // smallroom 에 불을 끈다 

	TurnLights(rooms)  // 한꺼번에 불을 켬 
}