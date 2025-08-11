package main

import "fmt"

func main() {
	// 4.1 다양한 타입의 변수 선언 (초기화하지 않음)
	// Go는 변수를 선언하고 초기화하지 않으면 자동으로 해당 타입의 "제로 값"으로 초기화한다.
	// Java에서는 지역 변수를 초기화하지 않으면 컴파일 에러가 발생한다.
	var a int     // int 타입의 제로 값은 0
	var b float64 // float64 타입의 제로 값은 0.0
	var c bool    // bool 타입의 제로 값은 false
	var d string  // string 타입의 제로 값은 "" (빈 문자열)
	var e *int    // 포인터 타입의 제로 값은 nil (Java의 null과 유사)

	// 4.2 제로 값 출력
	fmt.Println("int의 제로 값:", a)
	fmt.Println("float64의 제로 값:", b)
	fmt.Println("bool의 제로 값:", c)
	fmt.Println("string의 제로 값:", d)
	fmt.Println("포인터의 제로 값:", e)
}
