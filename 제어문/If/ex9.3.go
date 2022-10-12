package main

import "fmt"

func main() {
	age, err := fmt.Scan()

	if age > 30 {
		fmt.Println("you are 30s", err)
	} else if age < 30 && age>= 20 {
		fmt.Println("you are 20s, Young!")
	} else {
		fmt.Println("Are you teenager?")
	}
}