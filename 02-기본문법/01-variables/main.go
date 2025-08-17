package main

import "fmt"

func main() {
	// 1.1 `var` 키워드를 사용한 변수 선언 및 초기화
	// Java: `String characterName = "Hossam";`와 유사.
	// Go는 타입이 뒤에 온다.
	var characterName string = "Hossam"
	// Java: `int characterLevel = 99;`와 유사.
	var characterLevel int = 99

	// 1.2 `var` 키워드와 타입 추론을 사용한 변수 선언
	// 초기화 값으로 타입을 추론한다. Java 10+의 `var exp = 85.5;`와 유사.
	var exp = 85.5

	// 1.3 짧은 선언(`:=`)을 사용한 변수 선언 및 초기화
	// 함수 내부에서만 사용 가능하며, `var` 키워드와 타입 지정을 생략한다.
	// Go에서 가장 흔하게 사용되는 변수 선언 방식이다.
	isAlive := true

	// 1.4 변수 값 출력
	// `fmt.Println` 함수는 인자들을 공백으로 구분하여 출력하고 줄바꿈한다.
	fmt.Println("캐릭터 이름:", characterName)
	fmt.Println("레벨:", characterLevel)
	fmt.Println("경험치:", exp)
	fmt.Println("생존 여부:", isAlive)
}
