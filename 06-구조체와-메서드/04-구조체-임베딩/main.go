package main

import "fmt"

// 기본 캐릭터 구조체 (부모 역할)
type BaseCharacter struct {
	Name  string
	Level int
	HP    int
	MP    int
}

// 기본 캐릭터 메서드들
func (bc *BaseCharacter) DisplayBasicInfo() {
	fmt.Printf("이름: %s, 레벨: %d, HP: %d, MP: %d\n",
		bc.Name, bc.Level, bc.HP, bc.MP)
}

func (bc *BaseCharacter) LevelUp() {
	bc.Level++
	bc.HP += 10
	bc.MP += 5
	fmt.Printf("%s가 레벨업! (레벨: %d)\n", bc.Name, bc.Level)
}

// 전사 구조체 (임베딩을 통한 "상속")
type Warrior struct {
	BaseCharacter     // 익명 필드로 임베딩
	Strength      int // 전사만의 고유 필드
	Armor         int // 방어력
}

// 전사의 고유 메서드
func (w *Warrior) Attack() {
	damage := w.Strength * 2
	fmt.Printf("%s가 검으로 공격! 데미지: %d\n", w.Name, damage)
}

// 메서드 오버라이딩 (BaseCharacter의 LevelUp을 재정의)
func (w *Warrior) LevelUp() {
	w.BaseCharacter.LevelUp() // 부모의 메서드 호출
	w.Strength += 3           // 전사만의 추가 스탯 증가
	w.Armor += 2
	fmt.Printf("전사 스탯 증가: 힘 +3, 방어력 +2\n")
}

// 마법사 구조체
type Mage struct {
	BaseCharacter     // 임베딩
	Intelligence  int // 지능
	ManaRegen     int // 마나 재생력
}

// 마법사의 고유 메서드
func (m *Mage) CastSpell(spellName string) {
	manaCost := 20
	if m.MP >= manaCost {
		m.MP -= manaCost
		damage := m.Intelligence * 3
		fmt.Printf("%s가 %s 마법 시전! 데미지: %d (남은 MP: %d)\n",
			m.Name, spellName, damage, m.MP)
	} else {
		fmt.Printf("%s의 마나가 부족합니다!\n", m.Name)
	}
}

// 마법사의 레벨업 오버라이딩
func (m *Mage) LevelUp() {
	m.BaseCharacter.LevelUp()
	m.Intelligence += 4
	m.ManaRegen += 2
	fmt.Printf("마법사 스탯 증가: 지능 +4, 마나재생 +2\n")
}

// 전사 생성자
func NewWarrior(name string) *Warrior {
	return &Warrior{
		BaseCharacter: BaseCharacter{
			Name:  name,
			Level: 1,
			HP:    120,
			MP:    30,
		},
		Strength: 15,
		Armor:    10,
	}
}

// 마법사 생성자
func NewMage(name string) *Mage {
	return &Mage{
		BaseCharacter: BaseCharacter{
			Name:  name,
			Level: 1,
			HP:    80,
			MP:    100,
		},
		Intelligence: 20,
		ManaRegen:    5,
	}
}

func main() {
	fmt.Println("=== 구조체 임베딩 테스트 ===")

	// 캐릭터 생성
	warrior := NewWarrior("Thor")
	mage := NewMage("Gandalf")

	fmt.Println("초기 상태:")
	warrior.DisplayBasicInfo() // 임베딩된 메서드 자동 사용 가능
	mage.DisplayBasicInfo()

	fmt.Println("\n=== 고유 스킬 사용 ===")
	warrior.Attack()
	mage.CastSpell("파이어볼")

	fmt.Println("\n=== 레벨업 (오버라이딩된 메서드) ===")
	warrior.LevelUp()
	mage.LevelUp()

	fmt.Println("\n레벨업 후 상태:")
	warrior.DisplayBasicInfo()
	mage.DisplayBasicInfo()

	// 임베딩된 필드에 직접 접근도 가능
	fmt.Printf("\n전사의 힘: %d, 방어력: %d\n", warrior.Strength, warrior.Armor)
	fmt.Printf("마법사의 지능: %d, 마나재생: %d\n", mage.Intelligence, mage.ManaRegen)
}
