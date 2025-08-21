package main

import "fmt"

// GameCharacter 구조체를 정의함.
type GameCharacter struct {
	Name  string // 캐릭터 이름
	Level int    // 레벨
	HP    int    // 체력
	MP    int    // 마나
}

// '값 수신자(Value Receiver)' 메서드
// (c GameCharacter)는 이 메서드가 GameCharacter 타입에 속해 있음을 나타냄.
// 'c'는 메서드 내에서 구조체 인스턴스의 복사본에 접근할 때 사용하는 변수임.
// 원본 데이터를 수정하지 않고 단순히 정보를 보여주는 경우에 값 수신자를 사용함.
func (c GameCharacter) DisplayInfo() {
	fmt.Printf("이름: %s, 레벨: %d, HP: %d, MP: %d\n", c.Name, c.Level, c.HP, c.MP)
}

// '포인터 수신자(Pointer Receiver)' 메서드
// (c *GameCharacter)는 메서드가 구조체의 포인터에 바인딩됨을 의미함.
// Java에서 객체의 메서드가 'this'를 통해 원본 객체를 수정하는 것과 유사함.
// 메서드 내에서 구조체의 원본 필드 값을 변경해야 할 때 포인터 수신자를 사용함.
func (c *GameCharacter) LevelUp() {
	c.Level++   // c가 가리키는 원본 인스턴스의 Level 필드를 1 증가시킴.
	c.HP += 100 // 원본 인스턴스의 HP를 100 증가시킴.
	fmt.Printf("%s가 레벨업! (현재 레벨: %d)\n", c.Name, c.Level)
}

// 프로그램의 시작점인 main 함수임.
func main() {
	// GameCharacter 인스턴스를 생성하고 초기화함.
	character := GameCharacter{
		Name:  "Hossam",
		Level: 99,
		HP:    8000,
		MP:    12000,
	}

	// DisplayInfo 메서드를 호출함.
	// 이 메서드는 '값 수신자'이므로 character 인스턴스의 복사본을 사용하여 정보를 출력함.
	fmt.Println("--- 초기 상태 ---")
	character.DisplayInfo()

	// LevelUp 메서드를 호출함.
	// 이 메서드는 '포인터 수신자'이므로 character 인스턴스의 원본 데이터를 직접 수정함.
	fmt.Println("\n--- 레벨업 진행 ---")
	character.LevelUp()

	// 레벨업 후 변경된 상태를 확인하기 위해 다시 DisplayInfo 메서드를 호출함.
	// Level과 HP가 변경된 것을 확인할 수 있음.
	fmt.Println("\n--- 레벨업 후 상태 ---")
	character.DisplayInfo()
}
