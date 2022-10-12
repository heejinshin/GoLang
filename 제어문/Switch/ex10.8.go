// switch문_ const 열거값에 따라 수행되는 로직을 변경 
	// 색생값을 문자열로 바꾸기 

package main

import "fmt"

type ColorType int
const (
	CoralPink ColorType = iota
	PinkSky 
	MintChoco
)

func colorToString(color ColorType) string{
	switch color{
	case CoralPink:
		return "Coral"
	case PinkSky:
		return "Sky"
	case MintChoco:
		return "Mint"

	default:        // default는 생략가능 
		return "Undefined"
	}
}


func getMyFavoriteColor() ColorType{   
	return MintChoco
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}

