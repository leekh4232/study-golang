package main

import "fmt"

// 빈 인터페이스를 매개변수로 받는 함수.
// 어떤 타입의 값이든 이 함수에 전달할 수 있음.
func PrintAnything(v interface{}) {
	fmt.Printf("값: %v, 타입: %T\n", v, v)
}

func main() {
	fmt.Println("--- 빈 인터페이스 테스트 ---")

	// 정수형 값을 전달
	PrintAnything(10)

	// 문자열 값을 전달
	PrintAnything("안녕하세요")

	// 구조체 값을 전달
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Hossam", Age: 30}
	PrintAnything(p)
}
