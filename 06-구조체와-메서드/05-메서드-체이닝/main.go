package main

import "fmt"

// CharacterBuilder 구조체
type CharacterBuilder struct {
	character *GameCharacter
}

// GameCharacter 구조체 (기존 정의 재사용)
type GameCharacter struct {
	Name   string
	Level  int
	HP     int
	MP     int
	Class  string
	Skills []string
}

// 새로운 빌더 생성
func NewCharacterBuilder() *CharacterBuilder {
	return &CharacterBuilder{
		character: &GameCharacter{
			Level:  1,
			HP:     100,
			MP:     50,
			Skills: make([]string, 0),
		},
	}
}

// 체이닝 메서드들 (모두 *CharacterBuilder를 반환)
func (cb *CharacterBuilder) SetName(name string) *CharacterBuilder {
	cb.character.Name = name
	return cb // 자기 자신을 반환하여 체이닝 가능
}

func (cb *CharacterBuilder) SetClass(class string) *CharacterBuilder {
	cb.character.Class = class
	return cb
}

func (cb *CharacterBuilder) SetLevel(level int) *CharacterBuilder {
	cb.character.Level = level
	return cb
}

func (cb *CharacterBuilder) SetStats(hp, mp int) *CharacterBuilder {
	cb.character.HP = hp
	cb.character.MP = mp
	return cb
}

func (cb *CharacterBuilder) AddSkill(skill string) *CharacterBuilder {
	cb.character.Skills = append(cb.character.Skills, skill)
	return cb
}

// 최종 캐릭터 반환
func (cb *CharacterBuilder) Build() *GameCharacter {
	return cb.character
}

// 캐릭터 정보 출력 메서드
func (gc *GameCharacter) DisplayFullInfo() {
	fmt.Printf("=== %s (%s) ===\n", gc.Name, gc.Class)
	fmt.Printf("레벨: %d, HP: %d, MP: %d\n", gc.Level, gc.HP, gc.MP)
	fmt.Printf("스킬: %v\n", gc.Skills)
	fmt.Println()
}

func main() {
	fmt.Println("=== 메서드 체이닝으로 캐릭터 생성 ===")

	// 메서드 체이닝을 사용한 캐릭터 생성
	warrior := NewCharacterBuilder().
		SetName("Arthur").
		SetClass("기사").
		SetLevel(25).
		SetStats(300, 100).
		AddSkill("검술").
		AddSkill("방패막기").
		AddSkill("돌진").
		Build()

	mage := NewCharacterBuilder().
		SetName("Merlin").
		SetClass("대마법사").
		SetLevel(30).
		SetStats(200, 400).
		AddSkill("파이어볼").
		AddSkill("아이스스피어").
		AddSkill("라이트닝").
		AddSkill("힐링").
		Build()

	// 결과 출력
	warrior.DisplayFullInfo()
	mage.DisplayFullInfo()
}
