package main

import "fmt"

func main() {
	characterLevel := 99 // 캐릭터 레벨 변수 선언 및 초기화

	// 5.1 기본 if-else 문
	// Go의 `if` 조건문에는 소괄호 `()`를 사용하지 않는다. 중괄호 `{}`는 필수이다.
	// Java의 `if (characterLevel >= 100)`와 유사하다.
	if characterLevel >= 100 {
		fmt.Println("만렙 달성!")
	} else { // `else`는 `if`의 닫는 중괄호 `}`와 같은 줄에 와야 한다.
		fmt.Println("아직 만렙이 아님.")
	}

	// 5.2 `if` 문 내에서 변수 선언 및 사용 (if-short statement)
	// `if` 문 내에서만 유효한 변수를 선언하고 초기화할 수 있다. `exp` 변수는 `if` 블록 내에서만 사용 가능하다.
	// 이는 Java에는 없는 Go의 특징적인 문법이다.
	if exp := 85.5; exp > 90.0 {
		fmt.Println("경험치 충분, 레벨 업 가능!")
	} else {
		fmt.Println("경험치가 부족함.")
	}
}
