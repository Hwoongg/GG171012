package main

import "fmt"

func main() {
	//짧은 선언 : 선언과 동시에 초기화
	// 함수 안에서만 사용가능 (전역x), 선언 후 할당 예외 발생
	// 주로 제한된 범위의 함수내에서 사용할 경우 가독성 up
	// shortVar1 := 3
	// shortVar2 := "Test"
	// shortVar3 := false

	//var f, g, h int = 1, 2, 3

	// if안에서 사용
	if i := 10; i < 11 {
		fmt.Println("Short Var Test")
	}
}
