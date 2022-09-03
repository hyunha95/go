package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { // 매개변수의 타입이 같다면 뒤에 것만 적어주면 된다.
	return a * b
}

func lenAndUpper(name string) (int, string) { // 다른 타입 두개를 리턴할 수 있다.
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) { // 원하는 만큼 인자를 전달하는 방법
	fmt.Println(words)
}

func main() {
	name := "nico" // 축약형은 오로지 func안에서만 가능하고 변수에만 적용 가능하다.
	// var name string = "nico"  // 위 코드와 동일
	name = "Lynn"
	fmt.Println(name)

	fmt.Println(multiply(2, 2))

	totalLength, upperName := lenAndUpper("nico") // 변수명에 _를 줘서 무시하는 값으로 사용할 수 있다
	fmt.Println(totalLength, upperName)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}
