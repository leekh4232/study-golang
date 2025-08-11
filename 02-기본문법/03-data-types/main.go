package main

import "fmt"

func main() {
	// 3.1 정수 타입
	// `int`는 시스템 아키텍처에 따라 크기가 달라진다 (32비트 또는 64비트).
	var level int = 10
	// `int64`는 64비트 크기의 정수 타입이다. Java의 `long`과 유사하다.
	var monsterCount int64 = 100

	// Go는 타입에 매우 엄격하여 다른 타입 간의 직접적인 연산을 허용하지 않는다.
	// Java에서는 `int`와 `long` 간의 암시적 형 변환이 가능하지만, Go에서는 명시적으로 변환해야 한다.
	// fmt.Println(level + monsterCount) // 컴파일 에러 발생: `mismatched types int and int64`

	// 명시적 형 변환: `int64(level)`은 `level` 변수를 `int64` 타입으로 변환한다.
	fmt.Println("총합:", int64(level)+monsterCount)

	// 3.2 실수 타입
	// `float32`는 32비트 부동소수점 타입이다. Java의 `float`과 유사하다.
	var attackPower float32 = 35.5
	// `float64`는 64비트 부동소수점 타입이다. Java의 `double`과 유사하다.
	var defensePower float64 = 52.8

	// 다른 실수 타입 간에도 명시적 형 변환이 필수이다.
	fmt.Println("전투력:", float64(attackPower)*defensePower)

	// 3.3 문자열과 rune
	// `string` 타입은 UTF-8 인코딩된 문자열을 나타낸다. Java의 `String`과 유사하다.
	var message string = "Hello, World!"
	// `rune` 타입은 유니코드 코드 포인트(Unicode code point)를 나타내며, 실제로는 `int32`의 별칭이다.
	// 작은따옴표(`'`)를 사용하여 단일 문자를 표현한다. Java의 `char`와 유사하다.
	var firstChar rune = 'H'

	fmt.Println(message)
	// `rune`을 `string`으로 변환하여 출력한다. `string(firstChar)`는 `char`를 `String`으로 변환하는 것과 유사하다.
	fmt.Println("첫 글자:", string(firstChar))
}
