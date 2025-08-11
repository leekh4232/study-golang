package main

import "fmt"

func main() {
	state := "idle" // 캐릭터 상태 변수 선언

	// 7.1 기본 switch 문
	// Go의 `switch` 문은 Java의 `switch`와 유사하지만, `break`를 명시하지 않아도 각 `case`가 끝나면 자동으로 멈춘다.
	// Java의 `switch`는 `break`가 없으면 다음 `case`로 넘어간다 (fallthrough).
	switch state {
	case "idle": // `state`가 "idle"인 경우
		fmt.Println("캐릭터가 쉬고 있음.")
	case "walk", "run": // 여러 조건을 콤마(,)로 구분하여 한 번에 처리할 수 있다.
		fmt.Println("캐릭터가 이동 중임.")
	default: // 모든 `case`에 해당하지 않는 경우 (Java의 `default`와 동일)
		fmt.Println("알 수 없는 상태임.")
	}

	// 7.2 조건식을 사용한 switch (Java의 `if-else if` 체인과 유사)
	// `switch` 뒤에 변수 없이 `case`에 직접 조건식을 사용할 수 있다.
	level := 45 // 캐릭터 레벨 변수 선언
	switch {    // 조건식 없이 `switch` 키워드만 사용
	case level < 30: // `level`이 30 미만인 경우
		fmt.Println("초보자 레벨임.")
	case level >= 30 && level < 50: // `level`이 30 이상 50 미만인 경우
		fmt.Println("중급자 레벨임.")
	default: // 위 조건들에 모두 해당하지 않는 경우
		fmt.Println("고수 레벨임!")
	}
}
