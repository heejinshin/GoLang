package main 

import "fmt"

type Data struct {
	value int 
	data [200]int 
}

func ChangeData(arg Data) {  // arg 에 data가 복사가 됐다 (우변으로 동작했다  )
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data 

	ChangeData(data)   // arg로 들어갈때 사이즈가 복사됨, 공간은 다르다 
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

