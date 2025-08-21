package main

import "fmt"

// 빈 인터페이스 값을 받아 타입에 따라 다른 동작을 수행하는 함수
func CheckType(v interface{}) {
	fmt.Printf("입력 값: %v\n", v)

	// 1. "value, ok" 패턴을 사용한 안전한 타입 어설션
	// v를 string 타입으로 변환 시도
	str, ok := v.(string)
	if ok {
		fmt.Printf("  (ok) 이 값은 문자열입니다: %s\n", str)
	} else {
		fmt.Println("  (ok) 이 값은 문자열이 아닙니다.")
	}

	// 2. switch 문을 사용한 타입 어설션 (Type Switch)
	// 여러 타입을 검사할 때 훨씬 간결하고 효율적임.
	switch t := v.(type) {
	case int:
		fmt.Printf("  (switch) 이 값은 정수이며, 2를 곱하면 %d 입니다.\n", t*2)
	case string:
		fmt.Printf("  (switch) 이 값은 문자열이며, 길이는 %d 입니다.\n", len(t))
	case bool:
		fmt.Printf("  (switch) 이 값은 불리언입니다: %t\n", t)
	default:
		// t의 타입은 v의 타입과 동일
		fmt.Printf("  (switch) 처리할 수 없는 타입입니다: %T\n", t)
	}
	fmt.Println("--------------------")
}

func main() {
	CheckType("Hello, Go!")
	CheckType(123)
	CheckType(true)
	CheckType(3.14)
}
