package main

import "fmt"

// int 포인터를 인자로 받아 원본 int 값을 변경하는 함수
func incrementWithPointer(x *int) {
	*x++ // x가 가리키는 값(원본 변수)을 1 증가
	fmt.Println("  incrementWithPointer 함수 내부 *x:", *x)
}

// Player 구조체 포인터를 인자로 받아 원본 구조체 필드를 변경하는 함수
func levelUpWithPointer(p *Player) {
	p.Level++ // (*p).Level++ 와 동일 (Go가 자동으로 역참조)
	fmt.Println("  levelUpWithPointer 함수 내부 Player:", *p)
}

type Player struct {
	Name  string
	Level int
}

func main() {
	// 3.1 기본 타입에 포인터 사용
	fmt.Println("---", "3.1 기본 타입에 포인터 사용", "---")
	a := 5
	fmt.Println("main 함수 시작 a:", a) // 5
	incrementWithPointer(&a)        // a의 주소를 전달
	fmt.Println("main 함수 종료 a:", a) // 6 (변화 있음)

	// 3.2 구조체에 포인터 사용
	fmt.Println("\n---", "3.2 구조체에 포인터 사용", "---")
	p := Player{Name: "Warrior", Level: 10}
	fmt.Println("main 함수 시작 Player:", p) // {Warrior 10}
	levelUpWithPointer(&p)               // p의 주소를 전달
	fmt.Println("main 함수 종료 Player:", p) // {Warrior 11} (변화 있음)

	// 3.3 new 함수를 이용한 포인터 생성
	// new(T)는 T 타입의 제로 값으로 초기화된 변수를 할당하고 그 주소를 반환한다.
	newInt := new(int) // int 타입의 포인터 (*int) 반환, *newInt는 0으로 초기화
	fmt.Println("\n---", "3.3 new 함수를 이용한 포인터 생성", "---")
	fmt.Println("newInt 포인터 값:", newInt)    // 메모리 주소
	fmt.Println("newInt가 가리키는 값:", *newInt) // 0
	*newInt = 123
	fmt.Println("newInt가 가리키는 값 변경 후:", *newInt) // 123

	newPlayer := new(Player) // Player 타입의 포인터 (*Player) 반환, 필드는 제로 값으로 초기화
	fmt.Println("newPlayer 포인터 값:", newPlayer)
	fmt.Println("newPlayer가 가리키는 값:", *newPlayer) // { 0}
	newPlayer.Name = "Mage"
	newPlayer.Level = 5
	fmt.Println("newPlayer가 가리키는 값 변경 후:", *newPlayer) // {Mage 5}
}
