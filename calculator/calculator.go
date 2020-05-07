package main

import (
	"fmt"
)

func Plus(_x int32, _y int32) int32 {
	return _x + _y
}

func Minus(_x int32, _y int32) int32 {
	return _x - _y
}

func Multiply(_x int32, _y int32) int32 {
	return _x * _y
}

func main() {
	fmt.Print("중간고사 계산기 프로젝트입니다. GG171012황현웅\n")

	// Input Variable Group

	var (
		operator int32
		x        int32
		y        int32
	)

	// Input Part
	fmt.Print("계산을 원하는 두개의 변수를 입력해주세요.\n")
	fmt.Print("첫번째 변수 : ")
	fmt.Scan(&x)
	fmt.Print("두번째 변수 : ")
	fmt.Scan(&y)
	fmt.Print("연산을 골라주세요\n")
	fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈\n")
	fmt.Scan(&operator)

	switch operator {
	case 1:
		fmt.Println("덧셈 결과값 : ", Plus(x, y))
	case 2:
		fmt.Println("뺄셈 결과값 :", Minus(x, y))
	case 3:
		fmt.Println("곱셈 결과값 : ", Multiply(x, y))

	}

}
