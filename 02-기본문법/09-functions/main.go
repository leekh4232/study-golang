package main

import "fmt"

// 9.1 함수 정의 및 다중 반환 값
// `func` 키워드로 함수를 정의한다. 매개변수와 반환 타입이 명시된다.
// Go 함수는 여러 개의 반환 값을 가질 수 있다. (Java는 단일 값만 반환 가능)
// 입력: `name` (string), `level` (int), `exp` (float64)
// 반환: `summary` (string), `canLevelUp` (bool)
func getCharacterSummary(name string, level int, exp float64) (string, bool) {
	// 함수 내부에서 상수 선언
	const MAX_LEVEL = 100
	const EXP_REQUIRED_FOR_LEVEL_UP = 90.0

	// `fmt.Sprintf`는 포맷팅된 문자열을 반환한다. Java의 `String.format()`과 유사하다.
	summary := fmt.Sprintf("이름: %s, 레벨: %d, 경험치: %.1f", name, level, exp)
	// 레벨업 가능 여부를 판단하는 논리식
	canLevelUp := level < MAX_LEVEL && exp >= EXP_REQUIRED_FOR_LEVEL_UP

	// 두 개의 값을 반환한다.
	return summary, canLevelUp
}

func main() {
	// 9.2 함수 호출 및 다중 반환 값 받기
	// `summary`와 `canLevelUp` 변수에 각각 반환 값을 할당한다.
	summary, canLevelUp := getCharacterSummary("Hossam", 99, 85.5)

	fmt.Println(summary)

	// 9.3 반환된 `canLevelUp` 값을 이용한 조건문 처리
	if canLevelUp { // `canLevelUp`이 `true`이면 실행
		fmt.Println("레벨 업이 가능함!")
	} else { // `canLevelUp`이 `false`이면 실행
		fmt.Println("레벨 업까지 경험치가 더 필요함.")
	}

	// 9.4 다른 인자로 함수를 한 번 더 호출
	summary2, canLevelUp2 := getCharacterSummary("Alice", 80, 95.1)
	fmt.Println(summary2)
	if canLevelUp2 {
		fmt.Println("레벨 업이 가능함!")
	} else {
		fmt.Println("레벨 업까지 경험치가 더 필요함.")
	}
}
