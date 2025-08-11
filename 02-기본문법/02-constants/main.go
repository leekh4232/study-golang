package main

import "fmt"

// 2.1 단일 상수 선언
// `const` 키워드를 사용하여 상수를 선언한다. Java의 `final` 키워드와 유사하다.
const MAX_LEVEL = 100
const SERVER_IP = "127.0.0.1"

// 2.2 `iota`를 사용한 그룹 상수 선언
// `iota`는 0부터 시작하여 `const` 블록 내에서 선언된 상수마다 1씩 자동으로 증가하는 특별한 상수 생성기이다.
// Java의 `enum`과 유사하게 열거형 상수를 정의할 때 유용하다.
const (
	STATE_IDLE = iota // 첫 번째 `iota`는 0. STATE_IDLE = 0
	STATE_WALK        // `iota`가 1 증가하여 STATE_WALK = 1
	STATE_RUN         // `iota`가 1 증가하여 STATE_RUN = 2
	STATE_JUMP        // `iota`가 1 증가하여 STATE_JUMP = 3
)

func main() {
	// 2.3 상수 값 출력
	fmt.Println("최대 레벨:", MAX_LEVEL)
	fmt.Println("서버 IP:", SERVER_IP)
	fmt.Println("캐릭터 상태 - 대기:", STATE_IDLE)
	fmt.Println("캐릭터 상태 - 점프:", STATE_JUMP)
}
