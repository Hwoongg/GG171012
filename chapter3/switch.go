package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 제어문 - switch
	// 자동 break때문에 fallthrough 존재
	a := -7

	switch b := 27; {
	case b < 0:
		fmt.Println(a, "는 음수")

	case b == 0:
		fmt.Println(a, "는 0")
	case b > 0:
		fmt.Println(a, "는 양수")
	}

	rand.Seed(time.Now(), UnixNano())
	switch i := rand.Intn(100); {
	case i >= 50 && i < 100:
		fmt.Println("i->", i, " 50이상 100미만")
	}

	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a->", a, "짝수")
	case 1, 3, 5:
		fmt.Println("a->", a, "홀수")
	}
}
