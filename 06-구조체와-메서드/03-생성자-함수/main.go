package main

import "fmt"

// GameCharacter 구조체 정의
type GameCharacter struct {
	Name  string
	Level int
	HP    int
	MP    int
	Class string // 직업 필드 추가
}

// 기본 캐릭터 생성자 함수
// Go의 관례에 따라 'New' + 타입명으로 함수명을 정함
// 포인터를 반환하여 메모리 효율성을 높임
func NewGameCharacter(name, class string) *GameCharacter {
	return &GameCharacter{
		Name:  name,
		Level: 1,   // 기본 레벨 1로 시작
		HP:    100, // 기본 체력 100
		MP:    50,  // 기본 마나 50
		Class: class,
	}
}

// 고급 캐릭터 생성자 함수 (커스텀 스탯)
// Java의 생성자 오버로딩과 유사한 효과
func NewAdvancedCharacter(name, class string, level, hp, mp int) *GameCharacter {
	return &GameCharacter{
		Name:  name,
		Level: level,
		HP:    hp,
		MP:    mp,
		Class: class,
	}
}

// 팩토리 패턴: 직업별 캐릭터 생성
func NewWarrior(name string) *GameCharacter {
	return &GameCharacter{
		Name:  name,
		Level: 1,
		HP:    150, // 전사는 체력이 높음
		MP:    30,  // 마나는 낮음
		Class: "전사",
	}
}

func NewMage(name string) *GameCharacter {
	return &GameCharacter{
		Name:  name,
		Level: 1,
		HP:    80,  // 마법사는 체력이 낮음
		MP:    120, // 마나는 높음
		Class: "마법사",
	}
}

// 캐릭터 정보 표시 메서드 (값 수신자)
func (c GameCharacter) DisplayInfo() {
	fmt.Printf("이름: %s, 직업: %s, 레벨: %d, HP: %d, MP: %d\n",
		c.Name, c.Class, c.Level, c.HP, c.MP)
}

// 레벨업 메서드 (포인터 수신자)
func (c *GameCharacter) LevelUp() {
	c.Level++
	// 직업별로 다른 스탯 증가
	switch c.Class {
	case "전사":
		c.HP += 20
		c.MP += 5
	case "마법사":
		c.HP += 10
		c.MP += 15
	default:
		c.HP += 15
		c.MP += 10
	}
	fmt.Printf("%s(%s)가 레벨업! (현재 레벨: %d)\n", c.Name, c.Class, c.Level)
}

func main() {
	fmt.Println("=== 다양한 캐릭터 생성 방법 ===")

	// 1. 기본 생성자 사용
	character1 := NewGameCharacter("Alice", "궁수")
	fmt.Println("1. 기본 생성자로 생성:")
	character1.DisplayInfo()

	// 2. 고급 생성자 사용
	character2 := NewAdvancedCharacter("Bob", "도적", 10, 200, 80)
	fmt.Println("\n2. 고급 생성자로 생성:")
	character2.DisplayInfo()

	// 3. 팩토리 패턴 사용
	warrior := NewWarrior("Charlie")
	mage := NewMage("Diana")

	fmt.Println("\n3. 팩토리 패턴으로 생성:")
	warrior.DisplayInfo()
	mage.DisplayInfo()

	fmt.Println("\n=== 레벨업 테스트 ===")
	warrior.LevelUp()
	mage.LevelUp()

	fmt.Println("\n레벨업 후:")
	warrior.DisplayInfo()
	mage.DisplayInfo()
}
