package main

import "fmt"

func main() {
	var a int = 20
	b := 30

	if a >= 15 {
		fmt.Println("15 이상")
	}

	if b >= 25 {
		fmt.Println("25 이상")
	}

	if b := 30; b >= 25 {
		fmt.Println("25 이상")
	}
}
