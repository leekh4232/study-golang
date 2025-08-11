package main

import "fmt"

type Player struct {
	Name  string
	Level int
	HP    int
}

func main() {
	// 8.1 구조체 선언 및 초기화
	p1 := Player{Name: "Hero", Level: 10, HP: 100}
	fmt.Println("플레이어 이름:", p1.Name)

	// 8.2 필드 값 변경
	p1.Level = 11
	fmt.Println("플레이어 레벨:", p1.Level)
}
